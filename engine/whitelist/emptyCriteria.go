package whitelist

import (
	"errors"

	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/models"
)

func CreateEmptyCriteria(input model.NewEmptyCriteriaInput, authDetails *internal.AuthDetails) (*model.Drop, error) {

	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop")
	}

	drop, err := engine.GetDropByID(input.DropID)
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
		DropID:    drop.ID.String(),
		CreatorID: creator.ID.String(),
	}

	// emptyCriteria := model.ClaimCriteriaTypeEmptyCriteria
	// drop.Criteria = &emptyCriteria

	itemUpdateErr := engine.SaveModel(drop)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	criteriaUpdateErr := engine.SaveModel(criteria)
	if criteriaUpdateErr != nil {
		return nil, criteriaUpdateErr
	}

	return drop.ToGraphData(nil), nil
}
