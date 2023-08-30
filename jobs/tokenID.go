package jobs

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/mintwatcher"
	"inverse.so/models"
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

		if collection.ContractAddress == nil {
			continue
		}

		tokenID, err := fetchTokenUri(*collection.ContractAddress, item.ID.String())
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
	err := dbutils.DB.Where("token_id IS NULL").Find(&items).Error
	if err != nil {
		return nil, err
	}

	return &items, nil
}

func fetchTokenUri(contractAddress, itemID string) (*int, error) {

	inverseAPIBaseURL := "https://inverse-prod.onrender.com"
	rpcProvider := "https://polygon-mainnet.infura.io/v3/022bed77e57c4bcfae626a7f8bcf44de"
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
