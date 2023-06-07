package whitelist

import (
	"errors"

	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
	"inverse.so/services"
)

func CreateTwitterCriteria(input model.NewTwitterCriteriaInput, authDetails *internal.AuthDetails) (*model.Item, error) {

	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	tweetID, err := services.StripTweetIDFromLink(*input.TweetLink)
	if err != nil {
		return nil, errors.New("invalid tweet link")
	}

	var interactions string
	for i := range input.Interaction {
		interactions += string(*input.Interaction[i]) + ","
	}

	criteria := &models.TwitterCriteria{
		ItemID:       item.ID.String(),
		CreatorID:    creator.ID.String(),
		ProfileLink:  *input.ProfileLink,
		TweetLink:    *input.TweetLink,
		TweetID:      *tweetID,
		Interactions: interactions,
		CutOffDate:   *input.CutOffDate,
	}

	criteriaUpdateErr := engine.SaveModel(criteria)
	if criteriaUpdateErr != nil {
		return nil, criteriaUpdateErr
	}

	return item.ToGraphData(), nil
}
