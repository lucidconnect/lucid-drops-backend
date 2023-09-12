package claimers

import (
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

	mappedItems := make([]*model.Item, len(items))

	for idx, item := range items {
		mappedItems[idx] = item.ToGraphData()
	}

	return mappedItems, nil
}

func fetchAAAddressFromSignerInfo(address string) (*string, error) {

	var signerInfo models.SignerInfo
	err := dbutils.DB.Where("wallet_address = ?", address).First(&signerInfo).Error
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &signerInfo.WalletAddress, nil
}
