package whitelist

import (
	"errors"
	"fmt"

	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/models"
)

/** Criteria
- Members of a Farcaster Channel
- People who follow a specific user
- People who reply to, recast and like a specific cast
*/

func CreateFarcasterWhitelistForItem(input *model.NewFarcasterCriteriaInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop")
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

	if len(input.Interaction) == 0 {
		return nil, errors.New("please pass in some farcaster interaction type")
	}
	var interactions string
	for _, interaction := range input.Interaction {
		interactions += fmt.Sprintf("%v,", interaction.String())
	}

	criteria := &models.FarcasterCriteria{
		DropId:       item.DropID,
		CreatorID:    creator.ID,
		CastUrl:      *input.Cast,
		CriteriaType: input.CriteriaType,
	}
	item.Criteria = &input.CriteriaType
	if err = engine.SaveModel(item); err != nil {
		return nil, err
	}

	if err = engine.SaveModel(criteria); err != nil {
		return nil, err
	}
	
	return item.ToGraphData(), nil
}
