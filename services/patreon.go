package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"inverse.so/structure"
	"inverse.so/utils"
)

func FetchPatreonAccessToken(code *string) (*structure.PatreonAccessTokenResponse, error) {

	codeParams := url.Values{}
	codeParams.Set("code", *code)

	params := url.Values{}
	params.Set("client_id", utils.UseEnvOrDefault("PATREON_CLIENT_ID", ""))
	params.Set("client_secret", utils.UseEnvOrDefault("PATREON_CLIENT_SECRET", ""))
	params.Set("redirect_uri", utils.UseEnvOrDefault("PATREON_REDIRECT_URI", "https://5b62-102-216-201-35.ngrok-free.app/patreon/callback/"))

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
