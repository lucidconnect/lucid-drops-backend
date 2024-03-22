package whitelist

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/services/neynar"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
)

/** Criteria
- Members of a Farcaster Channel
- People who follow a specific user
- People who reply to, recast and like a specific cast


How do I handle a scenerio with a requirement such as:
- must be following xyz
- must interact with cast
- must be following chanel xyz

pass in multiple claim criterias
run a loop to get the passed in criteria

database:
- should allow a drop have multiple farcaster criteria (1:many relationship)
- migrate relationship  btw both tables from 1:1 to 1:many
- maybe create a table to hold the active criterias for a drop?

pass in an array of claim criteria type
store the array as a comma seperated string in db
during validation, parse the comma separated string into an array and validate each type in a loop
*/

func CreateFarcasterWhitelistForDrop(input model.NewFarcasterCriteriaInput, authDetails *internal.AuthDetails) (*model.Drop, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop")
	}

	drop, err := engine.GetDropByID(input.DropID)
	if err != nil {
		return nil, errors.New("drop not found")
	}

	if drop.Criteria != nil {
		//Delete Existing criteria
		err := engine.DeleteCriteriaIfExists(drop)
		if err != nil {
			return nil, err
		}
	}

	var interactions string
	for _, interaction := range input.Interaction {
		interactions += fmt.Sprintf("%v,", interaction.String())
	}
	var criteriaTypes string
	for _, criteriaType := range input.CriteriaType {
		criteriaTypes += fmt.Sprintf("%v,", criteriaType.String())
	}

	criteria := &models.FarcasterCriteria{
		DropId:       drop.ID,
		CreatorID:    creator.ID,
		CriteriaType: criteriaTypes,
	}
	if input.CastURL != nil {
		criteria.CastUrl = *input.CastURL
	}
	if input.ChannelID != nil {
		criteria.ChannelID = *input.ChannelID
	}
	if input.FarcasterUserName != nil {
		// resolve farcaster id by username
		apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
		neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
		if err != nil {
			log.Err(err).Send()
			return nil, err
		}

		fid, err := neynarClient.FetchFarcasterUserByUsername(*input.FarcasterUserName)
		if err != nil {
			return nil, err
		}
		criteria.FarcasterProfileID = fmt.Sprint(fid)
	}

	// drop.Criteria = &input.CriteriaType
	// if err = engine.SaveModel(drop); err != nil {
	// 	return nil, err
	// }

	if err = engine.SaveModel(criteria); err != nil {
		return nil, err
	}

	return drop.ToGraphData(nil), nil
}

func ValidateFarcasterCriteriaForDrop(farcasterAddress string, dropId string) (*model.ValidationRespoonse, error) {
	resp := &model.ValidationRespoonse{
		Valid: false,
	}

	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Send()
		return resp, err
	}

	userFid, err := neynarClient.FetchFarcasterUserFidByEthAddress(farcasterAddress)
	if err != nil {
		return resp, err
	}

	drop, err := engine.GetDropByID(dropId)
	if err != nil {
		return nil, err
	}

	if drop.FarcasterCriteria.ID == uuid.Nil {
		return nil, errors.New("drop does not have a farcaster criteria")
	}

	criteria := drop.FarcasterCriteria
	requiredCriteriaTypes := strings.Split(criteria.CriteriaType, ",")

	for _, criteriaType := range requiredCriteriaTypes {
		if criteriaType == model.ClaimCriteriaTypeFarcasterInteractions.String() {
			for _, interaction := range models.InteractionsToArr(criteria.Interactions) {
				switch *interaction {
				case model.InteractionTypeReplies:
					if !validateFarcasterReplyCriteria(int32(userFid), *criteria) {
						return resp, errors.New("farcaster account does not meet the reply criteria")
					}
				case model.InteractionTypeRecasts:
					if !validateFarcasterRecastCriteria(int32(userFid), *criteria) {
						return resp, errors.New("farcaster account does not meet the recast criteria")
					}
				case model.InteractionTypeLikes:
					if !validateFarcasterLikeCriteria(int32(userFid), *criteria) {
						return resp, errors.New("farcaster account does not meet the like criteria")
					}
				}
			}
		}

		if criteriaType == model.ClaimCriteriaTypeFarcasterFollowing.String() {
			if !validateFarcasterAccountFollowerCriteria(int32(userFid), *criteria) {
				return nil, errors.New("farcaster account does not meet the follower criteria")
			}
		}

		if criteriaType == model.ClaimCriteriaTypeFarcasterChannel.String() {
			if !validateFarcasterChannelFollowerCriteria(int32(userFid), *criteria) {
				return nil, errors.New("farcaster account does not meet the channel follower criteria")
			}
		}
	}
	passId, err := createMintPassForFarcasterMint(drop)
	if err != nil {
		return nil, err
	}

	resp.Valid = true
	resp.PassID = &passId
	return resp, nil
}
func validateFarcasterChannelFollowerCriteria(fid int32, criteria models.FarcasterCriteria) bool {
	var followers neynar.ChannelFollowers
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Send()
		return false
	}

	followers, err = neynarClient.RetrieveChannelFollowers(criteria.ChannelID, fid, "")
	if err != nil {
		return false
	}

	for followers.Next.Cursor != "" {
		for _, follower := range followers.Users {
			if follower.Fid == fid {
				return true
			}
		}
		followers, err = neynarClient.RetrieveChannelFollowers(criteria.ChannelID, fid, followers.Next.Cursor)
		if err != nil {
			return false
		}
	}

	return false
}

func validateFarcasterLikeCriteria(fid int32, criteria models.FarcasterCriteria) bool {
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Send()
		return false
	}

	cast, err := neynarClient.RetrieveCastByUrl(criteria.CastUrl)
	if err != nil {
		return false
	}

	for _, like := range cast.Reactions.Likes {
		if like.Fid == fid {
			return true
		}
	}

	return false
}

func validateFarcasterRecastCriteria(fid int32, criteria models.FarcasterCriteria) bool {
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Send()
		return false
	}

	cast, err := neynarClient.RetrieveCastByUrl(criteria.CastUrl)
	if err != nil {
		return false
	}

	for _, recast := range cast.Reactions.Recasts {
		if recast.Fid == fid {
			return true
		}
	}
	return false
}

func validateFarcasterReplyCriteria(fid int32, criteria models.FarcasterCriteria) bool {
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Send()
		return false
	}

	rootCast, err := neynarClient.RetrieveCastByUrl(criteria.CastUrl)
	if err != nil {
		log.Err(err).Send()
		return false
	}

	casts, err := neynarClient.RetrieveCastsByThreadHash(rootCast.Hash)
	if err != nil {
		log.Err(err).Send()
		return false
	}

	for _, cast := range casts {
		if cast.Authour.Fid == fid {
			return true
		}
	}
	return false
}

func validateFarcasterAccountFollowerCriteria(fid int32, criteria models.FarcasterCriteria) bool {
	var followers []neynar.RelevantFollowersDehydrated
	apiKeyOpt := neynar.WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := neynar.NewNeynarClient(apiKeyOpt)
	if err != nil {
		log.Err(err).Send()
		return false
	}

	creatorFid, err := strconv.Atoi(criteria.FarcasterProfileID)
	if err != nil {
		log.Err(err).Send()
		return false
	}
	followers, err = neynarClient.FetchRelvantFollowers(int32(creatorFid))
	if err != nil {
		return false
	}

	for _, follower := range followers {
		if follower.User.Fid == fid {
			return true
		}
	}
	return false
}

func createMintPassForFarcasterMint(drop *models.Drop) (string, error) {
	if drop.AAContractAddress == nil {
		return "", errors.New("drop contract address not found")
	}

	newMint := models.MintPass{
		DropID:              drop.ID.String(),
		DropContractAddress: *drop.AAContractAddress,
		BlockchainNetwork:   drop.BlockchainNetwork,
	}

	err := dbutils.DB.Create(&newMint).Error
	if err != nil {
		return "", err
	}

	return newMint.ID.String(), err
}
