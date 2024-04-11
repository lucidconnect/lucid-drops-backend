package route

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/lucidconnect/inverse/drops"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/mintwatcher"
	"github.com/lucidconnect/inverse/services/neynar"
	"github.com/lucidconnect/inverse/utils"
	"github.com/lucidconnect/inverse/whitelist"
	"github.com/rs/zerolog/log"
)

type Server struct {
	HttpServer *http.Server
	router     *chi.Mux
	nftRepo    drops.NFTRepository
}

func NewServer(port string, nftRepo drops.NFTRepository, router *chi.Mux) *Server {
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	return &Server{
		HttpServer: httpServer,
		router:     router,
		nftRepo:    nftRepo,
	}
}
func (s *Server) CreateMintPass(w http.ResponseWriter, r *http.Request) {
	var err error
	pass := &model.ValidationRespoonse{
		Valid: false,
	}

	dropId := r.URL.Query().Get("dropId")
	walletAddress := r.URL.Query().Get("wallet")
	drop, _ := s.nftRepo.FindDropById(dropId)

	if drop.EditionLimit != nil {
		count, err := s.nftRepo.CountMintPassesForDrop(dropId)
		if err != nil {
			log.Err(err).Caller().Send()
			json.NewEncoder(w).Encode(pass)
			return
		}
		if int(count) >= *drop.EditionLimit {
			pass.Message = utils.GetStrPtr("this nft has reached it's mint")
			log.Err(err).Caller().Send()
			json.NewEncoder(w).Encode(pass)
			return
		}
	}

	if drop.FarcasterCriteria != nil {
		apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
		neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
		if err != nil {
			log.Err(err).Caller().Send()
			return
		}

		pass, err = neynarClient.ValidateFarcasterCriteriaForDrop(walletAddress, *drop)
		if err != nil {
			log.Err(err).Caller().Msg("GenerateSignatureForClaim")
			json.NewEncoder(w).Encode(pass)
			return
		}
	}

	mintPass, err := s.nftRepo.GetMintPassForWallet(dropId, walletAddress)
	if err == nil {
		if drop.UserLimit != nil {
			passes, err := s.nftRepo.CountMintPassesForAddress(dropId, walletAddress)
			if err == nil {
				if int(passes) >= *drop.UserLimit {
					pass.Valid = false
					pass.Message = utils.GetStrPtr("limit reached for wallet")
					json.NewEncoder(w).Encode(pass)
					return
				}
			}
		}
	} else {
		mintPass = &drops.MintPass{
			DropID:              dropId,
			DropContractAddress: *drop.AAContractAddress,
			BlockchainNetwork:   drop.BlockchainNetwork,
			MinterAddress:       walletAddress,
			TokenID:             "1",
		}
		if err = s.nftRepo.CreateMintPass(mintPass); err != nil {
			log.Err(err).Caller().Send()
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	pass.PassID = utils.GetStrPtr(mintPass.ID.String())
	pass.Valid = true
	if err = json.NewEncoder(w).Encode(pass); err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) GenerateSignatureForClaim(w http.ResponseWriter, r *http.Request) {
	passId := r.URL.Query().Get("passId")
	claimingAddress := r.URL.Query().Get("claimingAddress")
	input := model.GenerateClaimSignatureInput{
		OtpRequestID:    passId,
		ClaimingAddress: claimingAddress,
	}

	now := time.Now()
	mintPass, err := s.nftRepo.GetMintPassById(input.OtpRequestID)
	if err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if mintPass.UsedAt != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// passes, err := s.nftRepo.CountMintPassesForAddress(mintPass.DropID, input.ClaimingAddress)
	// if err == nil {
	// 	if passes != 0 {
	// 		err = errors.New("more than one mint pass found for this minter address")
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 	}
	// }

	mintPass.UsedAt = &now
	err = s.nftRepo.UpdateMintPass(mintPass)
	if err != nil {
		log.Err(err).Caller().Send()
		return
	}

	sig, err := whitelist.GenerateSignatureForClaim(*mintPass)
	if err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(sig); err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// this method isn't really needed right now
func (s *Server) VerifyItemTokenIDs() {

	items, err := s.nftRepo.FindItemsWithUnresolvesTokenIDs()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	for _, item := range items {
		drop, err := s.nftRepo.FindDropById(item.DropID.String())
		if err != nil {
			log.Error().Msg(err.Error())
			continue
		}

		if drop.AAContractAddress == nil {
			continue
		}

		var isBase bool
		if drop.BlockchainNetwork != nil {
			isBase = *drop.BlockchainNetwork == model.BlockchainNetworkBase
		}

		tokenID, err := fetchTokenUri(*drop.AAContractAddress, item.ID.String(), isBase)
		if err != nil {
			log.Error().Msg(err.Error())
			continue
		}

		if tokenID == nil {
			log.Info().Msgf("🚨 Token ID not found for Item %s", item.ID)
			continue
		}

		tokenIDint64 := int64(*tokenID)
		item.TokenID = tokenIDint64

		// err = engine.SaveModel(&item)
		// if err != nil {
		// 	log.Error().Msg(err.Error())
		// }
	}
}

func fetchTokenUri(contractAddress, itemID string, isBase bool) (*int, error) {
	inverseAPIBaseURL := os.Getenv("API_BASE_URL")
	rpcProvider := utils.UseEnvOrDefault("POLYGON_RPC_PROVIDER", "https://polygon-mainnet.g.alchemy.com/v2/wH3GkDxLOS4h8O7hmIPWqvmOvE4VIqWn")
	if isBase {
		rpcProvider = utils.UseEnvOrDefault("BASE_RPC_PROVIDER", "https://base-mainnet.g.alchemy.com/v2/2jx1c05x5vFN7Swv9R_ZJKKAXZUfas8A")
	}

	client, err := ethclient.Dial(rpcProvider)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	addressToAddress := common.HexToAddress(contractAddress)
	x, err := mintwatcher.NewMintwatcher(addressToAddress, client)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	opts := &bind.CallOpts{}
	// TODO make counter more dynamic
	for i := 1; i <= 10; i++ {
		expectedURI := fmt.Sprintf("%s/metadata/%s/%s", inverseAPIBaseURL, contractAddress, itemID)
		integer := big.NewInt(int64(i))
		uri, err := x.Uri(opts, integer)
		if err != nil {
			log.Error().Msg(err.Error())
			log.Info().Msgf("🔍 Fetching token URI for %s/%s", contractAddress, itemID)
		}

		if uri == expectedURI {
			return &i, nil
		}
	}
	return nil, errors.New("token ID not found")
}
