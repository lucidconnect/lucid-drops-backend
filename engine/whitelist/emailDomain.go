package whitelist

import (
	"errors"

	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
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

	err = engine.SaveModel(&models.EmailDomainWhiteList{
		ItemID:     item.ID,
		CreatorID:  creator.ID,
		BaseDomain: input.AuthorizedSubdomain,
	})

	if err != nil {
		return nil, err
	}

	emailCriteria := model.ClaimCriteriaTypeEmailDomain
	item.Criteria = &emailCriteria
	itemUpdateErr := engine.SaveModel(item)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	return item.ToGraphData(), nil
}
