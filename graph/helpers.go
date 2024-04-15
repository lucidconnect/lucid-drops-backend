package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lucidconnect/inverse/drops"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/services"
	"github.com/lucidconnect/inverse/services/neynar"
	"github.com/lucidconnect/inverse/utils"
	"github.com/rs/zerolog/log"
)

func createMintUrl(item, imagUrl, contract string) (string, error) {
	baseurl := os.Getenv("FRAME_SERVER")
	url := fmt.Sprintf("%v/createframe", baseurl)

	type createFrameRequest struct {
		DropId     string `json:"dropId"`
		ImageUrl   string `json:"imageUrl"`
		Collection string `json:"collection"`
	}

	request := createFrameRequest{
		DropId:     item,
		ImageUrl:   imagUrl,
		Collection: contract,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	httpRequest, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	httpRequest.Header.Set("Content-Type", "appication/json")

	res, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var frameUrl string
	if err = json.NewDecoder(res.Body).Decode(&frameUrl); err != nil {
		return "", err
	}

	return frameUrl, nil
}

func walletLimitReached(walletAddress string, pass drops.MintPass) bool {
	// set default claim limit to 1
	// var mintsByAddress int64
	// var pass drops.MintPass
	// err := dbutils.DB.Model(&drops.MintPass{}).Where("drop_id = ?", dropID).Where("minter_address = ?", walletAddress).First(&pass).Error
	// if err != nil {
	// 	return false
	// } else {
	var alchemyOpts []services.Option
	apiKeyOpt := services.WithApiKey(os.Getenv("ALCHEMY_API_KEY"))
	urlOpt := services.WithUrl(os.Getenv("ALCHEMY_URL"))
	alchemyOpts = append(alchemyOpts, apiKeyOpt, urlOpt)
	alchemyClient, err := services.NewAlchemyClient(alchemyOpts...)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	holders, err := alchemyClient.GetOwnersForNft(pass.DropContractAddress, "1")
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	for _, address := range holders {
		if address == walletAddress {
			return true
		}
	}
	// }

	return false
}

// func dropOverEditionLimit(drop drops.Drop) bool {
// 	if drop.EditionLimit != nil {
// 		var editionCount int64
// 		err := dbutils.DB.Model(&models.MintPass{}).Where("drop_id = ?", drop.ID).Count(&editionCount).Error
// 		if err == nil {
// 			return int(editionCount) >= *drop.EditionLimit
// 		}
// 	}
// 	return false
// }

// func validateFarcasterCriteriaForDrop(farcasterAddress string, drop drops.Drop) (*model.ValidationRespoonse, error) {
// 	resp := &model.ValidationRespoonse{
// 		Valid: false,
// 	}

// 	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
// 	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
// 	if err != nil {
// 		log.Err(err).Send()
// 		return nil, err
// 	}

// 	userFid, err := neynarClient.FetchFarcasterUserFidByEthAddress(farcasterAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	criteria := drop.FarcasterCriteria
// 	requiredCriteriaTypes := strings.Split(criteria.CriteriaType, ",")

// 	for _, criteriaType := range requiredCriteriaTypes {
// 		if criteriaType == model.ClaimCriteriaTypeFarcasterInteractions.String() {
// 			for _, interaction := range models.InteractionsToArr(criteria.Interactions) {
// 				switch *interaction {
// 				case model.InteractionTypeReplies:
// 					if !validateFarcasterReplyCriteria(int32(userFid), *criteria) {
// 						return resp, errors.New("farcaster account does not meet the reply criteria")
// 					}
// 				case model.InteractionTypeRecasts:
// 					if !validateFarcasterRecastCriteria(int32(userFid), *criteria) {
// 						return resp, errors.New("farcaster account does not meet the recast criteria")
// 					}
// 				case model.InteractionTypeLikes:
// 					if !validateFarcasterLikeCriteria(int32(userFid), *criteria) {
// 						return resp, errors.New("farcaster account does not meet the like criteria")
// 					}
// 				}
// 			}
// 		}

// 		if criteriaType == model.ClaimCriteriaTypeFarcasterFollowing.String() {
// 			if !validateFarcasterAccountFollowerCriteria(int32(userFid), *criteria) {
// 				return resp, errors.New("farcaster account does not meet the follower criteria")
// 			}
// 		}

// 		if criteriaType == model.ClaimCriteriaTypeFarcasterChannel.String() {
// 			if !validateFarcasterChannelFollowerCriteria(int32(userFid), *criteria) {
// 				return resp, errors.New("farcaster account does not meet the channel follower criteria")
// 			}
// 		}
// 	}

// 	return resp, nil
// }

func validateFarcasterChannelFollowerCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	var followers neynar.ChannelFollowers
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	followers, err = neynarClient.RetrieveChannelFollowers(criteria.ChannelID, fid, "")
	if err != nil {
		return false
	}

	for followers.Next.Cursor != "" {
		for _, follower := range followers.Users {
			if follower.Fid == fid {
				return true
			}
		}
		followers, err = neynarClient.RetrieveChannelFollowers(criteria.ChannelID, fid, followers.Next.Cursor)
		if err != nil {
			return false
		}
	}

	return false
}

func validateFarcasterLikeCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	cast, err := neynarClient.RetrieveCastByUrl(criteria.CastUrl)
	if err != nil {
		return false
	}

	for _, like := range cast.Reactions.Likes {
		if like.Fid == fid {
			return true
		}
	}

	return false
}

func validateFarcasterRecastCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	cast, err := neynarClient.RetrieveCastByUrl(criteria.CastUrl)
	if err != nil {
		return false
	}

	for _, recast := range cast.Reactions.Recasts {
		if recast.Fid == fid {
			return true
		}
	}
	return false
}

func validateFarcasterReplyCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	rootCast, err := neynarClient.RetrieveCastByUrl(criteria.CastUrl)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	casts, err := neynarClient.RetrieveCastsByThreadHash(rootCast.Hash)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	for _, cast := range casts {
		if cast.Author.Fid == fid {
			return true
		}
	}
	return false
}

func validateFarcasterAccountFollowerCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	var followers []neynar.RelevantFollowersDehydrated
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	creatorFid, err := strconv.Atoi(criteria.FarcasterProfileID)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}
	followers, err = neynarClient.FetchRelvantFollowers(int32(creatorFid))
	if err != nil {
		return false
	}

	for _, follower := range followers {
		if follower.User.Fid == fid {
			return true
		}
	}
	return false
}

func fetchNftHolders(item *model.Item) ([]string, error) {
	var alchemyOpts []services.Option
	apiKeyOpt := services.WithApiKey(os.Getenv("ALCHEMY_API_KEY"))
	urlOpt := services.WithUrl(os.Getenv("ALCHEMY_URL"))
	alchemyOpts = append(alchemyOpts, apiKeyOpt, urlOpt)
	alchemyClient, err := services.NewAlchemyClient(alchemyOpts...)
	if err != nil {
		log.Err(err).Caller().Send()
		return nil, err
	}

	holders, err := alchemyClient.GetOwnersForNft(item.DropAddress, "1")
	if err != nil {
		log.Err(err).Caller().Send()
		return nil, err
	}
	return holders, nil
}

// func StoreUserAccountSignerAddress(input model.SignerInfo, authDetails *internal.AuthDetails) (bool, error) {
// 	creator, err := engine.GetCreatorByAddress(authDetails.Address)
// 	if err != nil {
// 		return false, err
// 	}

// 	noSignature := "NONE" // TODO use aa wallet signatures to authorize third party signer
// 	if input.Signature == nil {
// 		input.Signature = &noSignature
// 	}

// 	aaWallet := common.HexToAddress(input.Address)
// 	altSigner, err := engine.GetAltSignerByCreatorID(creator.ID.String())
// 	if err != nil {
// 		altSigner = &models.SignerInfo{
// 			CreatorID:     creator.ID.String(),
// 			WalletAddress: aaWallet.String(),
// 			Signature:     input.Signature,
// 			Provider:      input.Provider,
// 		}
// 	} else {
// 		altSigner.WalletAddress = aaWallet.String()
// 		altSigner.Provider = input.Provider
// 		altSigner.Signature = input.Signature
// 	}

// 	alterr := engine.SaveModel(altSigner)
// 	if alterr != nil {
// 		return false, err
// 	}

//		return true, nil
//	}
func IsThisAValidEthAddress(maybeAddress string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	if len(maybeAddress) != 43 {
		return false
	}

	return re.MatchString(maybeAddress)
}

func getEthBackend(rpc string) *ethclient.Client {
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		log.Err(err).Caller().Msg("Failed to connect to the Ethereum client")
	}
	return conn
}

func getChainId(network model.BlockchainNetwork) *big.Int {
	var chain *big.Int
	switch network {
	case model.BlockchainNetworkBase:
		if ok, _ := utils.IsProduction(); ok {
			chain = big.NewInt(8453)
		} else {
			chain = big.NewInt(84532)
		}
	default:
		return nil
	}
	return chain
}
