package whitelist

import (
	"errors"

	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/models"
)

func CreateEmptyCriteria(input model.NewEmptyCriteriaInput, authDetails *internal.AuthDetails) (*model.Item, error) {

	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	// if item.Criteria != nil {
	// 	//Delete Existing criteria
	// 	err := engine.DeleteCriteriaIfExists(item)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

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
