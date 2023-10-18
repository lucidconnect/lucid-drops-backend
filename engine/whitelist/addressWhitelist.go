package whitelist

import (
	"errors"
	"time"

	"gorm.io/gorm/clause"
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
)

func CreateWalletAddressWhitelistForItem(input *model.NewWalletAddressWhitelistInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new collection")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if item.Criteria != nil {
		//Delete Existing criteria
		err := engine.DeleteCriteriaIfExists(item)
		if err != nil {
			return nil, err
		}
	}

	if len(input.AuthorizedWalletAddresses) == 0 {
		return nil, errors.New("please passing in a wallet address list")
	}

	dbWallets := make([]*models.WalletAddressClaim, len(input.AuthorizedWalletAddresses))
	for idx, address := range input.AuthorizedWalletAddresses {
		dbWallets[idx] = &models.WalletAddressClaim{
			CreatorID:     creator.ID,
			ItemID:        item.ID,
			WalletAddress: address,
		}
	}

	insertionErr := dbutils.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dbWallets, 100).Error
	if insertionErr != nil {
		return nil, insertionErr
	}

	for _, dbWallet := range dbWallets {

		timeNow := time.Now()

		dbWallet.SentOutAt = &timeNow

		insertionErr := dbutils.DB.Save(dbWallet).Error

		if insertionErr != nil {
			continue
		}
	}

	addressCriteria := model.ClaimCriteriaTypeWalletAddress
	item.Criteria = &addressCriteria
	itemUpdateErr := engine.SaveModel(item)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	return item.ToGraphData(), nil
}
