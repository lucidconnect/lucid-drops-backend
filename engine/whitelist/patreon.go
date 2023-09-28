package whitelist

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
	"inverse.so/services"
	"inverse.so/structure"
)

func ProcessPatreonCallback(code *string, creator bool) (*string, []*structure.PatreonCampaignInfo, error) {

	patreonToken, err := services.FetchPatreonAccessToken(code, creator)
	if err != nil {
		return nil, nil, err
	}

	patreonDetails := &models.PatreonAuthDetails{
		Code:         *code,
		AccessToken:  patreonToken.AccessToken,
		ExpiresAt:    time.Now().Add(time.Second * time.Duration(patreonToken.ExpiresIn)),
		RefreshToken: patreonToken.RefreshToken,
	}

	user, err := services.FetchPatreonUserLocal(patreonDetails)
	if err != nil {
		return nil, nil, err
	}

	campaigns, err := services.FetchPatreonCampaignLocal(patreonDetails)
	if err != nil {
		return nil, nil, err
	}

	patreonDetails.UserID = user.Id
	jsonValues, _ := json.Marshal(user.MembershirUIDs)
	patreonDetails.MembershipUIDs = string(jsonValues)

	if len(campaigns) == 1 {
		patreonDetails.CampaignID = campaigns[0].Id
	}

	err = engine.SaveModel(patreonDetails)
	if err != nil {
		return nil, nil, err
	}

	ID := patreonDetails.ID.String()
	return &ID, campaigns, nil
}

func CreatePatreonCriteria(input model.NewPatreonCriteriaInput, authDetails *internal.AuthDetails) (*model.Item, error) {

	log.Print(authDetails.Address)
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

	criteria := &models.PatreonCriteria{
		ItemID:    item.ID.String(),
		CreatorID: creator.ID.String(),
		AuthID:    input.AuthID,
	}

	authdetails, err := engine.FetchPatreonAuthByID(input.AuthID)
	if err != nil {
		return nil, err
	}

	if input.CampaignID != nil {

		authdetails.CampaignID = *input.CampaignID
		err = engine.SaveModel(authdetails)
		if err != nil {
			return nil, err
		}
	}

	if input.CampaignName != nil {
		criteria.CampaignName = *input.CampaignName
	}

	criteriaUpdateErr := engine.SaveModel(criteria)
	if criteriaUpdateErr != nil {
		return nil, criteriaUpdateErr
	}

	patreonCriteria := model.ClaimCriteriaTypePatreon
	item.Criteria = &patreonCriteria
	err = engine.SaveModel(item)
	if err != nil {
		return nil, err
	}

	return item.ToGraphData(), nil
}

func ValidatePatreonCriteriaForItem(itemID string, authID *string) (*model.ValidationRespoonse, error) {

	resp := &model.ValidationRespoonse{
		Valid: false,
	}

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return resp, errors.New("item not found")
	}

	if item.ClaimDeadline != nil {
		if time.Now().After(*item.ClaimDeadline) {
			return nil, errors.New("the item is no longer available to be claimed")
		}
	}

	if item.PatreonCriteria == nil {
		return resp, errors.New("item does not have a patreon criteria")
	}

	if authID == nil {
		return resp, errors.New("no patreon auth id provided")
	}

	patreonAuth, err := engine.FetchPatreonAuthByID(*authID)
	if err != nil {
		return resp, errors.New("patreon account not authorized")
	}

	creatorAuth, err := engine.FetchPatreonAuthByID(item.PatreonCriteria.AuthID)
	if err != nil {
		return resp, errors.New("creator patreon account not authorized")
	}

	campaignPledges, err := services.FetchPatreonPledgesLocal(creatorAuth)
	if err != nil {
		return resp, errors.New("error fetching pledges")
	}

	var membershipIDs = make(map[string]string)
	_ = json.Unmarshal([]byte(patreonAuth.MembershipUIDs), &membershipIDs)
	if membershipIDs == nil {
		return resp, errors.New("user is not a valid patron")
	}

	for _, membershipID := range membershipIDs {

		_, ok := campaignPledges[membershipID]
		if ok {
			patreonAuth.WhiteListed = true
			err = engine.SaveModel(patreonAuth)
			if err != nil {
				return resp, errors.New("error saving patreon auth")
			}

			PassID, err := createMintPassForPatreonMint(item)
			if err != nil {
				return resp, errors.New("error creating mint pass")
			}

			resp.Valid = true
			resp.PassID = PassID
			return resp, nil
		}
	}

	return nil, errors.New("user is not a valid patron")
}

func createMintPassForPatreonMint(item *models.Item) (*string, error) {
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
	}

	err = dbutils.DB.Create(&newMint).Error
	if err != nil {
		return nil, err
	}

	passId := newMint.ID.String()
	return &passId, nil
}
