package onboarding

import (
	"errors"

	"gorm.io/gorm"
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/models"
)

func CreateCreatorProfileIfAddressIsMissing(address string) (*models.Creator, error) {
	cachedCreator, err := engine.GetCreatorByAddress(address)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newCreator := models.Creator{WalletAddress: address}
		creationErr := dbutils.DB.Create(&newCreator).Error
		if creationErr != nil {
			return nil, creationErr
		}
		return &newCreator, nil
	}

	return cachedCreator, nil
}
