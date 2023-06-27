package whitelist

import (
	"errors"

	"gorm.io/gorm/clause"
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
)

func CreateEmailWhitelistForItem(input *model.NewEmailWhitelistInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if len(input.AuthorizedEmails) == 0 {
		return nil, errors.New("please passing in email list")
	}

	dbEmails := make([]*models.SingleEmailClaim, len(input.AuthorizedEmails))
	for idx, email := range input.AuthorizedEmails {
		dbEmails[idx] = &models.SingleEmailClaim{
			CreatorID:    creator.ID,
			ItemID:       item.ID,
			EmailAddress: email,
		}
	}

	insertionErr := dbutils.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dbEmails, 100).Error
	if insertionErr != nil {
		return nil, insertionErr
	}

	emailCriteria := model.ClaimCriteriaTypeEmailWhiteList
	item.Criteria = &emailCriteria
	itemUpdateErr := engine.SaveModel(item)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	return item.ToGraphData(), nil
}
