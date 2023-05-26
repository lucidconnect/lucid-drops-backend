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

func GetCreatorByInverseUsername(inverseUsername string) (*models.Creator, error) {
	var creator models.Creator

	err := utils.DB.Where(&models.Creator{InverseUsername: &inverseUsername}).First(&creator).Error
	if err != nil {
		return nil, errors.New("username isn't being used")
	}

	return &creator, nil
}

func CreateCollection(newCollection *models.Collection) error {
	return utils.DB.Create(newCollection).Error
}
