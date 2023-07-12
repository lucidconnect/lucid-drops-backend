package whitelist

import (
	"inverse.so/engine"
	"inverse.so/models"
	"inverse.so/services"
)

func ProcessPatreonCallback(code *string) (*string, error) {

	patreonToken, err := services.FetchPatreonAccessToken(code)
	if err != nil {
		return nil, err
	}

	patreonDetails := &models.PatreonAuthDetails{
		Code:        *code,
		AccessToken: patreonToken.AccessToken,
	}

	// do other claim specific stuff here
	err = engine.SaveModel(patreonDetails)
	if err != nil {
		return nil, err
	}

	ID := patreonDetails.ID.String()
	return &ID, nil
}
