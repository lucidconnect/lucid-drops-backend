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

	var tweetLink string
	if input.TweetLink != nil {
		tweetLink = *input.TweetLink
	}

	var profileLink string
	if input.ProfileLink != nil {
		profileLink = *input.ProfileLink
	}

	var date string
	if input.CutOffDate != nil {
		date = *input.CutOffDate
	}

	criteria := &models.TwitterCriteria{
		ItemID:       item.ID.String(),
		CreatorID:    creator.ID.String(),
		ProfileLink:  profileLink,
		TweetLink:    tweetLink,
		TweetID:      *tweetID,
		CriteriaType: input.CriteriaType,
		Interactions: interactions,
		CutOffDate:   date,
	}

	twitterCriteria := input.CriteriaType
	item.Criteria = &twitterCriteria
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

func ProcessTwitterCallback(token, verifier *string) (*string, error) {

	tweetInfo, err := services.FetchTwitterAccessToken(token, verifier)
	if err != nil {
		return nil, err
	}

	// do other claim specific stuff here
	err = engine.SaveModel(tweetInfo)
	if err != nil {
		return nil, err
	}

	ID := tweetInfo.ID.String()
	return &ID, nil
}
