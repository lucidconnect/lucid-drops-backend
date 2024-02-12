package onboarding

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/models"
	"gorm.io/gorm"
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

			go CreateWalletAccount(newCreator.ID.String())
			return &newCreator, nil
		}
	}

	return cachedCreator, nil
}

func CreateWalletAccount(creatorID string) error {
	return dbutils.DB.Create(&models.Wallet{
		CreatorID: creatorID,
	}).Error
}
