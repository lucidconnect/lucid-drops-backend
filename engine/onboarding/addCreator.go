package onboarding

import (
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/models"
)

func CreateCreatorProfileIfAddressIsMissing(address string) (*models.Creator, error) {
	cachedCreator, err := engine.GetCreatorByAddress(address)
	if err != nil {
		newCreator := models.Creator{WalletAddress: address}
		creationErr := dbutils.DB.Create(&newCreator).Error
		if creationErr != nil {
			return nil, creationErr
		}
		return &newCreator, nil
	}

	return cachedCreator, nil
}
