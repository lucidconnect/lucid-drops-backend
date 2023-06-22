package whitelist

import (
	"errors"
	"strings"

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

	if item.TwitterCriteria != nil {
		return nil, errors.New("item already has a twitter criteria")
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

func ValidateTwitterCriteriaForItem(itemID, authID string) (bool, error) {

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return false, errors.New("item not found")
	}

	if item.TwitterCriteria == nil {
		return false, errors.New("item does not have a twitter criteria")
	}

	auth, err := engine.FetchTwitterAuthByID(authID)
	if err != nil {
		return false, errors.New("twitter account not authorized")
	}

	return validateTwitterAuthWithCriteria(auth, item.TwitterCriteria)
}

func validateTwitterAuthWithCriteria(auth *models.TwitterAuthDetails, criteria *models.TwitterCriteria) (bool, error) {

	// check reply criteria
	if criteria.CriteriaType == model.ClaimCriteriaTypeTwitterInteractions {

		for _, x := range models.InteractionsToArr(criteria.Interactions) {
			switch *x {
			case model.InteractionTypeReplies:
				if !validateReplyCriteria(auth, criteria) {
					return false, errors.New("twitter account does not meet the reply criteria")
				}
			case model.InteractionTypeRetweets:
				if !validateRetweetCriteria(auth, criteria) {
					return false, errors.New("twitter account does not meet the retweet criteria")
				}
			case model.InteractionTypeLikes:
				if !validateLikeCriteria(auth, criteria) {
					return false, errors.New("twitter account does not meet the like criteria")
				}
			}
		}

	}

	auth.WhiteListed = true
	auth.ItemID = &criteria.ItemID
	err := engine.SaveModel(auth)
	if err != nil {
		return false, err
	}

	return true, nil
}

func validateReplyCriteria(auth *models.TwitterAuthDetails, criteria *models.TwitterCriteria) bool {

	tweet, err := services.FetchTweetByID(criteria.TweetID)
	if err != nil {
		return false
	}

	for _, x := range tweet.Data.ThreadedConversationWithInjectionsV2.Instructions[0].Entries {

		// check if tweet is a reply
		if strings.Split(x.EntryID, "-")[0] != "conversationthread" {
			continue
		}

		for _, y := range x.Content.Items {
			if y.Item.ItemContent.TweetResults.Result.Core.UserResults.Result.RestID == auth.UserID {

				replyDeets := y.Item.ItemContent.TweetResults.Result.Legacy
				// check if tweet is a reply
				if replyDeets.InReplyToStatusIDStr == "" {
					return false
				}

				// check if tweet is a reply to the criteria tweet
				if replyDeets.InReplyToStatusIDStr != criteria.TweetID {
					return false
				}

				return true
			}
		}
	}

	return false
}

func validateRetweetCriteria(auth *models.TwitterAuthDetails, criteria *models.TwitterCriteria) bool {

	return false
}

func validateLikeCriteria(auth *models.TwitterAuthDetails, criteria *models.TwitterCriteria) bool {

	return false
}
