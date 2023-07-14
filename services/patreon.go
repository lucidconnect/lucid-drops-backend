package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

const (
	patreonAPIBaseURL = "https://www.patreon.com/api/oauth2/v2/"
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

func FetchPatreonUserLocal(auth *models.PatreonAuthDetails) (*structure.PatreonUserResponse, error) {

	err := refreshAccessTokenIfExpired(auth)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: auth.AccessToken,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	url := fmt.Sprintf("%sidentity?include=memberships&fields%%5Buser%%5D=full_name&fields%%5Btier%%5D=amount_cents", patreonAPIBaseURL)

	var user *patreonAuth.UserResponse
	err = executePatreonOAuthRequest(url, tc, &user)
	if err != nil {
		return nil, err
	}

	return &structure.PatreonUserResponse{
		Id:   user.Data.ID,
		Name: user.Data.Attributes.FullName,
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

func FetchPatreonCampaignLocal(auth *models.PatreonAuthDetails) ([]*structure.PatreonCampaignInfo, error) {

	err := refreshAccessTokenIfExpired(auth)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: auth.AccessToken,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	url := fmt.Sprintf("%scampaigns%s", patreonAPIBaseURL, "?fields%5Bcampaign%5D=creation_name")

	var campaign *structure.PatreonCampaigns
	err = executePatreonOAuthRequest(url, tc, &campaign)
	if err != nil {
		return nil, err
	}

	log.Printf("campaign info 🚨 %+v", campaign)
	var campaigns []*structure.PatreonCampaignInfo
	for _, campaign := range campaign.Data {

		singleCampaign, err := FetchCampaignLocal(auth, campaign.ID)
		if err != nil {
			return nil, err
		}

		campaigns = append(campaigns, &structure.PatreonCampaignInfo{
			Id:   campaign.ID,
			Name: singleCampaign.Name,
		})
	}

	return campaigns, nil
}

func FetchCampaignLocal(auth *models.PatreonAuthDetails, campaignID string) (*structure.PatreonCampaignInfo, error) {

	err := refreshAccessTokenIfExpired(auth)
	if err != nil {
		return nil, err
	}

	param := url.Values{}
	param.Set("fields[campaign]", "creation_name")
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: auth.AccessToken,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	url := fmt.Sprintf("%scampaigns/%s?%s", patreonAPIBaseURL, campaignID, param.Encode())

	var campaign *structure.PatreonCampaign
	err = executePatreonOAuthRequest(url, tc, &campaign)
	if err != nil {
		return nil, err
	}

	log.Printf("single campaign info 🚨 %+v", campaign)
	return &structure.PatreonCampaignInfo{
		Id:   campaign.Data.ID,
		Name: campaign.Data.Attributes.CreationName,
	}, nil
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

func FetchPatreonPledgesLocal(auth *models.PatreonAuthDetails) (map[string]*patreonAuth.UserResponse, error) {

	err := refreshAccessTokenIfExpired(auth)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: auth.AccessToken,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	url := fmt.Sprintf("%scampaigns/%s/members", patreonAPIBaseURL, auth.CampaignID)

	var pledge *structure.PatreonCampaignMembers
	users := make(map[string]*patreonAuth.UserResponse)
	for {

		err = executePatreonOAuthRequest(url, tc, &pledge)
		if err != nil {
			return nil, err
		}

		// u, ok := item.(*patreonAuth.UserResponse)
		// if !ok {
		// 	continue
		// }

		// users[u.Data.ID] = u

		if pledge.Meta.Pagination.Cursors.Next == "" {
			break
		}

		url = fmt.Sprintf("%s&cursor=%s", url, pledge.Meta.Pagination.Cursors.Next)

	}

	return users, nil
}

func executePatreonOAuthRequest(reqUrl string, client *http.Client, destination interface{}) error {

	//debugUtil
	// client.Transport = &http.Transport{
	// 	Proxy: func(req *http.Request) (*url.URL, error) {
	// 		return url.Parse("http://192.168.100.167:9090") //this sshould be dynamic based on the proxyman url
	// 	},
	// }

	log.Print("🚨 executing patreon request 🚨", reqUrl)
	resp, err := client.Get(reqUrl)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return json.Unmarshal(body, destination)
}

func refreshAccessTokenIfExpired(auth *models.PatreonAuthDetails) error {

	if auth.ExpiresAt.After(time.Now()) {
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
