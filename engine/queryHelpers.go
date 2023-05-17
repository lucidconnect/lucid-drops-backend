package engine

import (
	"errors"

	"inverse.so/models"
	"inverse.so/utils"
)

func GetCreatorByAddress(address string) (*models.Creator, error) {
	var creator models.Creator

	err := utils.DB.Where(&models.Creator{WalletAddress: address}).First(&creator).Error
	if err != nil {
		return nil, errors.New("address not found")
	}

	return &creator, nil
}

func CreateCreatorProfileIfAddressIsMissing(address string) (*models.Creator, error) {
	cachedCreator, err := GetCreatorByAddress(address)
	if err != nil {
		newCreator := models.Creator{WalletAddress: address}
		creationErr := utils.DB.Create(&newCreator).Error
		if creationErr != nil {
			return nil, creationErr
		}
		return &newCreator, nil
	}

	return cachedCreator, nil
}
