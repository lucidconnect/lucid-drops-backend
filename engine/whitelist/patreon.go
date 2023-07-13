package whitelist

import (
	"errors"
	"time"

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

	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
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

func ValidatePatreonCriteriaForItem(itemID string, authID *string) (bool, error) {

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return false, errors.New("item not found")
	}

	if item.PatreonCriteria == nil {
		return false, errors.New("item does not have a patreon criteria")
	}

	if authID == nil {
		return false, errors.New("no patreon auth id provided")
	}

	patreonAuth, err := engine.FetchPatreonAuthByID(*authID)
	if err != nil {
		return false, errors.New("patreon account not authorized")
	}

	creatorAuth, err := engine.FetchPatreonAuthByID(item.PatreonCriteria.AuthID)
	if err != nil {
		return false, errors.New("creator patreon account not authorized")
	}

	campaignPledges, err := services.FetchPledges(creatorAuth)
	if err != nil {
		return false, errors.New("error fetching pledges")
	}

	_, found := campaignPledges[patreonAuth.UserID]
	if !found {
		return false, errors.New("user not found in campaign pledges")
	}

	patreonAuth.WhiteListed = true
	err = engine.SaveModel(patreonAuth)
	if err != nil {
		return false, errors.New("error saving patreon auth")
	}

	return true, nil
}
