package claimers

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/models"
)

func FetchClaimedItems(address string) ([]*model.Item, error) {

	aaAddress, err := fetchAAAddressFromSignerInfo(address)
	if err == nil {
		address = *aaAddress
	}

	items, err := engine.GetClaimedItemByAddress(address)
	if err != nil {
		return nil, err
	}

	eoaClaimedItems, err := engine.GetClaimedItemsByEOAAddress(address)
	if err != nil {
		return nil, err
	}

	items = append(items, eoaClaimedItems...)
	mappedItems := make([]*model.Item, len(items))

	for idx, item := range items {
		mappedItems[idx] = item.ToGraphData()
	}

	return mappedItems, nil
}

func fetchAAAddressFromSignerInfo(address string) (*string, error) {

	creator, err := engine.GetCreatorByAddress(common.HexToAddress(address)) 
	if err != nil {
		return nil, err
	}

	var signerInfo models.SignerInfo
	err = dbutils.DB.Model(&models.SignerInfo{}).Where("creator_id = ?", creator.ID).First(&signerInfo).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &signerInfo.WalletAddress, nil
}
