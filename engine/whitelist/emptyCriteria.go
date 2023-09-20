package whitelist

import (
	"errors"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
	
)

func CreateEmptyCriteria(input model.NewEmptyCriteriaInput, authDetails *internal.AuthDetails) (*model.Item, error) {

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

	criteria := &models.EmptyCriteria{
		ItemID:    item.ID.String(),
		CreatorID: creator.ID.String(),
		
	}

	emptyCriteria := model.ClaimCriteriaTypeEmptyCriteria
	item.Criteria = &emptyCriteria

	itemUpdateErr := engine.SaveModel(item)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	criteriaUpdateErr := engine.SaveModel(criteria)
	if criteriaUpdateErr != nil {
		return nil, criteriaUpdateErr
	}

	return item.ToGraphData(), nil
}