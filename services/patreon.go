package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
	patreonAuth "gopkg.in/mxpv/patreon-go.v1"
	"inverse.so/engine"
	"inverse.so/models"
	"inverse.so/structure"
	"inverse.so/utils"
)

func FetchPatreonAccessToken(code *string, creator bool) (*structure.PatreonAccessTokenResponse, error) {

	codeParams := url.Values{}
	codeParams.Set("code", *code)

	params := url.Values{}
	params.Set("client_id", utils.UseEnvOrDefault("PATREON_CLIENT_ID", ""))
	params.Set("client_secret", utils.UseEnvOrDefault("PATREON_CLIENT_SECRET", ""))
	
	if creator {
		params.Set("redirect_uri", utils.UseEnvOrDefault("PATREON_CREATOR_REDIRECT_URI", "https://5b62-102-216-201-35.ngrok-free.app/patreon/callback/"))
	} else {
		params.Set("redirect_uri", utils.UseEnvOrDefault("PATREON_REDIRECT_URI", "https://5b62-102-216-201-35.ngrok-free.app/patreon/whitelist/callback/"))
	}

	url := fmt.Sprintf("https://www.patreon.com/api/oauth2/token?%s&grant_type=authorization_code&%s", codeParams.Encode(), params.Encode())
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var res *http.Response
	res, err = utils.GetDebugClient().Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var accessTokenResponse structure.PatreonAccessTokenResponse
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return nil, err
	}

	return &accessTokenResponse, nil
}

func RefreshPatreonAccessToken(refreshToken *string) (*structure.PatreonAccessTokenResponse, error) {

	params := url.Values{}
	params.Set("refresh_token", *refreshToken)
	params.Set("client_id", utils.UseEnvOrDefault("PATREON_CLIENT_ID", ""))
	params.Set("client_secret", utils.UseEnvOrDefault("PATREON_CLIENT_SECRET", ""))

	url := fmt.Sprintf("https://www.patreon.com/api/oauth2/token?grant_type=refresh_token&%s", params.Encode())
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var res *http.Response
	res, err = utils.GetDebugClient().Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var accessTokenResponse structure.PatreonAccessTokenResponse
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return nil, err
	}

	return &accessTokenResponse, nil
}

func FetchPatreonUser(auth *models.PatreonAuthDetails) (*structure.PatreonUserResponse, error) {

	err := refreshAccessTokenIfExpired(auth)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: auth.AccessToken,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	client := patreonAuth.NewClient(tc)

	userResponse, err := client.FetchUser()
	if err != nil {
		return nil, err
	}

	return &structure.PatreonUserResponse{
		Id:   userResponse.Data.ID,
		Name: userResponse.Data.Attributes.FullName,
	}, nil
}

func FetchCampaigns(auth *models.PatreonAuthDetails) ([]*structure.PatreonCampaignInfo, error) {

	err := refreshAccessTokenIfExpired(auth)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: auth.AccessToken,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	client := patreonAuth.NewClient(tc)

	campaignResponse, err := client.FetchCampaign()
	if err != nil {
		return nil, err
	}

	var campaigns []*structure.PatreonCampaignInfo
	for _, campaign := range campaignResponse.Data {
		campaigns = append(campaigns, &structure.PatreonCampaignInfo{
			Id:   campaign.ID,
			Name: campaign.Attributes.CreationName,
		})
	}

	return campaigns, nil
}

func FetchPledges(auth *models.PatreonAuthDetails) (map[string]*patreonAuth.UserResponse, error) {

	err := refreshAccessTokenIfExpired(auth)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: auth.AccessToken,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	client := patreonAuth.NewClient(tc)

	cursor := ""
	// var pledges []string
	users := make(map[string]*patreonAuth.UserResponse)
	for {
		pledgeResponse, err := client.FetchPledges(auth.CampaignID, patreonAuth.WithPageSize(100), patreonAuth.WithCursor(cursor))
		if err != nil {
			return nil, err
		}

		for _, item := range pledgeResponse.Included.Items {

			u, ok := item.(*patreonAuth.UserResponse)
			if !ok {
				continue
			}

			users[u.Data.ID] = u
			// pledges = append(pledges, u.Data.ID)
		}

		if pledgeResponse.Links.Next == "" {
			break
		}

		cursor = pledgeResponse.Links.Next
	}

	return users, nil
}

func refreshAccessTokenIfExpired(auth *models.PatreonAuthDetails) error {

	if auth.ExpiresAt.Before(time.Now()) {
		return nil
	}

	accessTokenResponse, err := RefreshPatreonAccessToken(&auth.RefreshToken)
	if err != nil {
		return err
	}

	auth.AccessToken = accessTokenResponse.AccessToken
	auth.RefreshToken = accessTokenResponse.RefreshToken
	auth.ExpiresAt = time.Now().Add(time.Second * time.Duration(accessTokenResponse.ExpiresIn))

	err = engine.SaveModel(auth)
	if err != nil {
		return err
	}

	return nil
}
