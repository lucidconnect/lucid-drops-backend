package whitelist

import (
	"encoding/json"
	"errors"
	"strings"

	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
	"inverse.so/services"
	"inverse.so/structure"
)

func CreateTwitterCriteria(input model.NewTwitterCriteriaInput, authDetails *internal.AuthDetails) (*model.Item, error) {

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

	var interactions string
	for i := range input.Interaction {
		interactions += string(*input.Interaction[i]) + ","
	}

	var tweetLink string
	var tweetID *string
	if input.TweetLink != nil {
		tweetID, err = services.StripTweetIDFromLink(*input.TweetLink)
		if err != nil {
			return nil, errors.New("invalid tweet link")
		}
		tweetLink = *input.TweetLink
	}

	var profileID string
	if input.ProfileID != nil {
		profileID = *input.ProfileID
	}

	var date string
	if input.CutOffDate != nil {
		date = *input.CutOffDate
	}

	criteria := &models.TwitterCriteria{
		ItemID:       item.ID.String(),
		CreatorID:    creator.ID.String(),
		ProfileID:    profileID,
		TweetLink:    tweetLink,
		TweetID:      *tweetID,
		CriteriaType: input.CriteriaType,
		Interactions: interactions,
		AuthID:       input.TwitterAuthID,
		CutOffDate:   date,
	}

	twitterCriteria := input.CriteriaType
	item.Criteria = &twitterCriteria
	itemUpdateErr := engine.SaveModel(nil, item)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	criteriaUpdateErr := engine.SaveModel(nil, criteria)
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
	err = engine.SaveModel(nil, tweetInfo)
	if err != nil {
		return nil, err
	}

	ID := tweetInfo.ID.String()
	return &ID, nil
}

func indexTweetRetweets(criteria models.TwitterCriteria) error {

	var retweeeterIDs []string
	details, err := engine.FetchTwitterAuthByID(criteria.AuthID)
	if err != nil {
		return err
	}

	retweets, err := services.FetchTweetRetweetsWithUserAuth(details, criteria.TweetID, nil)
	if err != nil {
		return err
	}

	retweeeterIDs = append(retweeeterIDs, retweets.Ids...)
	for retweets.NextCursorStr != "0" {

		retweets, err = services.FetchTweetRetweetsWithUserAuth(details, criteria.TweetID, &retweets.NextCursorStr)
		if err != nil {
			return err
		}

		mergeSliceOfStrings(retweeeterIDs, retweets.Ids)
	}

	retweetsToBytes, err := json.Marshal(retweeeterIDs)
	if err != nil {
		return err
	}

	criteria.IndexedRetweets = string(retweetsToBytes)
	err = engine.SaveModel(nil, criteria)
	if err != nil {
		return err
	}

	return nil
}

func indexTwitterFollowers(criteria models.TwitterCriteria) error {

	var followerIDs []string
	details, err := engine.FetchTwitterAuthByID(criteria.AuthID)
	if err != nil {
		return err
	}

	followers, err := services.FetchTwitterFollowersWithUserAuth(details, criteria.ProfileID, nil)
	if err != nil {
		return err
	}

	followerIDs = append(followerIDs, followers.Ids...)
	for followers.NextCursorStr != "0" {

		followers, err = services.FetchTwitterFollowersWithUserAuth(details, criteria.ProfileID, &followers.NextCursorStr)
		if err != nil {
			return err
		}

		mergeSliceOfStrings(followerIDs, followers.Ids)
	}

	followersToBytes, err := json.Marshal(followers)
	if err != nil {
		return err
	}

	criteria.IndexedFollowers = string(followersToBytes)
	err = engine.SaveModel(nil, criteria)
	if err != nil {
		return err
	}

	return nil
}

func ValidateTwitterCriteriaForItem(itemID, authID string) (*model.ValidationRespoonse, error) {

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if item.TwitterCriteria == nil {
		return nil, errors.New("item does not have a twitter criteria")
	}

	auth, err := engine.FetchTwitterAuthByID(authID)
	if err != nil {
		return nil, errors.New("twitter account not authorized")
	}

	return validateTwitterAuthWithCriteria(auth, item.TwitterCriteria, item)
}

func validateTwitterAuthWithCriteria(auth *models.TwitterAuthDetails, criteria *models.TwitterCriteria, Item *models.Item) (*model.ValidationRespoonse, error) {

	resp := &model.ValidationRespoonse{
		Valid: false,
	}

	// check reply criteria
	if criteria.CriteriaType == model.ClaimCriteriaTypeTwitterInteractions {

		for _, x := range models.InteractionsToArr(criteria.Interactions) {
			switch *x {
			case model.InteractionTypeReplies:
				if !validateReplyCriteria(auth, criteria) {
					return resp, errors.New("twitter account does not meet the reply criteria")
				}
			case model.InteractionTypeRetweets:
				if !validateRetweetCriteria(auth, criteria) {
					return resp, errors.New("twitter account does not meet the retweet criteria")
				}
			case model.InteractionTypeLikes:
				if !validateLikeCriteria(auth, criteria) {
					return resp, errors.New("twitter account does not meet the like criteria")
				}
			}
		}

	}

	if criteria.CriteriaType == model.ClaimCriteriaTypeTwitterFollowers {
		if !validateFollowerCriteria(auth, criteria) {
			return nil, errors.New("twitter account does not meet the follower criteria")
		}
	}

	PassID, err := createMintPassForTwitterMint(Item)
	if err != nil {
		return nil, errors.New("error creating mint pass")
	}

	resp.Valid = true
	resp.PassID = PassID

	auth.WhiteListed = true
	auth.ItemID = &criteria.ItemID
	err = engine.SaveModel(nil, auth)
	if err != nil {
		return resp, err
	}

	return resp, nil
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

	var retweetedBy *structure.TweetRetweetsResponse
	var err error
	retweetedBy, err = services.FetchTweetRetweets(criteria.TweetID, nil)
	if err != nil {
		return false
	}

	for retweetedBy.Meta.NextToken != "" {

		for _, x := range retweetedBy.Data {

			if x.ID == auth.UserID {
				return true
			}
		}

		retweetedBy, err = services.FetchTweetRetweets(criteria.TweetID, &retweetedBy.Meta.NextToken)
		if err != nil {
			return false
		}
	}

	return false
}

func validateLikeCriteria(auth *models.TwitterAuthDetails, criteria *models.TwitterCriteria) bool {

	var likingUsers *structure.TweetLikesResponse
	var err error
	likingUsers, err = services.FetchTweetLikingUsers(criteria.TweetID, nil)
	if err != nil {
		return false
	}

	for likingUsers.Meta.NextToken != "" {

		for _, x := range likingUsers.Data {

			if x.ID == auth.UserID {
				return true
			}
		}

		likingUsers, err = services.FetchTweetLikingUsers(criteria.TweetID, &likingUsers.Meta.NextToken)
		if err != nil {
			return false
		}
	}

	return false
}

func validateFollowerCriteria(auth *models.TwitterAuthDetails, criteria *models.TwitterCriteria) bool {

	var followers *structure.TwitterFollowersResponse
	var err error
	followers, err = services.FetchTwitterFollowers(criteria.ProfileID, nil)
	if err != nil {
		return false
	}

	for followers.Meta.NextToken != "" {

		for _, x := range followers.Data {

			if x.ID == auth.UserID {
				return true
			}
		}

		followers, err = services.FetchTwitterFollowers(criteria.ProfileLink, &followers.Meta.NextToken)
		if err != nil {
			return false
		}
	}

	return false
}

func createMintPassForTwitterMint(item *models.Item) (*string, error) {
	collection, err := engine.GetCollectionByID(item.CollectionID.String())
	if err != nil {
		return nil, errors.New("collection not found")
	}

	if collection.AAContractAddress == nil {
		return nil, errors.New("collection contract address not found")
	}

	if item.TokenID == nil {
		return nil, errors.New("The requested item is not ready to be claimed, please try again in a few minutes")
	}

	newMint := models.MintPass{
		ItemId:                    item.ID.String(),
		ItemIdOnContract:          *item.TokenID,
		CollectionContractAddress: *collection.AAContractAddress,
		BlockchainNetwork:         collection.BlockchainNetwork,
	}

	err = dbutils.DB.Create(&newMint).Error
	if err != nil {
		return nil, err
	}

	passId := newMint.ID.String()
	return &passId, nil
}

func mergeSliceOfStrings(s1, s2 []string) []string {

	return append(s1, s2...)
}
