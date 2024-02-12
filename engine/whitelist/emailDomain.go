package whitelist

import (
	"errors"

	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/models"
	"gorm.io/gorm/clause"
)

func CreateEmailDomainWhitelist(input *model.NewEmailDomainWhitelistInput, authDetails *internal.AuthDetails) (*model.Item, error) {
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

	dbEmails := make([]*models.EmailDomainWhiteList, len(input.AuthorizedSubdomains))

	for idx, domain := range input.AuthorizedSubdomains {
		dbEmails[idx] = &models.EmailDomainWhiteList{
			ItemID:     item.ID,
			CreatorID:  creator.ID,
			BaseDomain: domain,
		}
	}

	insertionErr := dbutils.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dbEmails, 100).Error
	if insertionErr != nil {
		return nil, insertionErr
	}

	emailCriteria := model.ClaimCriteriaTypeEmailDomain
	item.Criteria = &emailCriteria
	item.ShowEmailDomainHints = input.Visible

	itemUpdateErr := engine.SaveModel(item)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	return item.ToGraphData(), nil
}
