package jobs

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/mintwatcher"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/utils"
	"github.com/rs/zerolog/log"
)

func VerifyItemTokenIDs() {

	items, err := fetchItemsWithUnresolvedTokenIDs()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	for _, item := range *items {
		collection, err := engine.GetCollectionByID(item.CollectionID.String())
		if err != nil {
			log.Error().Msg(err.Error())
			continue
		}

		if collection.AAContractAddress == nil {
			continue
		}

		var isBase bool
		if collection.BlockchainNetwork != nil {
			isBase = *collection.BlockchainNetwork == model.BlockchainNetworkBase
		}

		tokenID, err := FetchTokenUri(*collection.AAContractAddress, item.ID.String(), isBase)
		if err != nil {
			log.Error().Msg(err.Error())
			continue
		}

		if tokenID == nil {
			log.Info().Msgf("🚨 Token ID not found for Item %s", item.ID)
			continue
		}

		tokenIDint64 := int64(*tokenID)
		item.TokenID = &tokenIDint64
		err = engine.SaveModel(&item)
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}
}

func fetchItemsWithUnresolvedTokenIDs() (*[]models.Item, error) {
	var items []models.Item
	oneHourAgo := time.Now().Add(-1 * time.Hour)
	err := dbutils.DB.Where("token_id IS NULL and created_at BETWEEN ? AND ?", oneHourAgo, time.Now()).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return &items, nil
}

func FetchTokenUri(contractAddress, itemID string, isBase bool) (*int, error) {
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
