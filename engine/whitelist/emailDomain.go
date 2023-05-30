package whitelist

import (
	"errors"

	"gorm.io/gorm/clause"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
	"inverse.so/utils"
)

func CreateEmailDomainWhitelist(input *model.NewEmailDomainWhitelistInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	dbEmails := make([]*models.EmailDomainWhiteList, len(input.AuthorizedSubdomains))

	for idx, domain := range input.AuthorizedSubdomains {
		dbEmails[idx] = &models.EmailDomainWhiteList{
			ItemID:     item.ID,
			CreatorID:  creator.ID,
			BaseDomain: domain,
		}
	}

	insertionErr := utils.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dbEmails, 100).Error
	if insertionErr != nil {
		return nil, insertionErr
	}

	emailCriteria := model.ClaimCriteriaTypeEmailDomain
	item.Criteria = &emailCriteria
	itemUpdateErr := engine.SaveModel(item)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	return item.ToGraphData(), nil
}
