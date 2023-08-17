package onboarding

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/models"
)

func CreateCreatorProfileIfAddressIsMissing(address common.Address) (*models.Creator, error) {
	cachedCreator, err := engine.GetCreatorByAddress(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newCreator := models.Creator{WalletAddress: address.String()}
			creationErr := dbutils.DB.Create(&newCreator).Error
			if creationErr != nil {
				return nil, creationErr
			}
			return &newCreator, nil
		}
	}

	return cachedCreator, nil
}
