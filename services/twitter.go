package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"inverse.so/graph/model"
	"inverse.so/models"
	"inverse.so/structure"
	"inverse.so/utils"
)

const (
	twitterAuthEntryPoint = "https://twitter.com/damndeji/status/1421901257988575234?s=21&t=-nCCU8hHFcvpr103Q9MWRQ"
	twwitterAPIURL        = "https://api.twitter.com/2/"
)

func FetchTweetLikingUsers(tweetID string, nextToken *string) (*structure.TweetLikesResponse, error) {

	endpoint := fmt.Sprintf("tweets/%s/liking_users", tweetID)
	if nextToken != nil {
		endpoint = fmt.Sprintf("%s?pagination_token=%s", endpoint, *nextToken)
	}

	var response structure.TweetLikesResponse
	err := executeTwitterRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func FetchTweetRetweets(tweetID string, nextToken *string) (*structure.TweetRetweetsResponse, error) {
	endpoint := fmt.Sprintf("tweets/%s/retweeted_by", tweetID)
	if nextToken != nil {
		endpoint = fmt.Sprintf("%s?pagination_token=%s", endpoint, *nextToken)
	}

	var response structure.TweetRetweetsResponse
	err := executeTwitterRequest("GET", endpoint, nil, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func FetchTwitterFollowers(profileID string, nextToken *string) (*structure.TwitterFollowersResponse, error) {
	endpoint := fmt.Sprintf("users/%s/followers", profileID)
	if nextToken != nil {
		endpoint = fmt.Sprintf("%s?pagination_token=%s", endpoint, *nextToken)
	}

	var response structure.TwitterFollowersResponse
	err := executeTwitterRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func FetchTweetRetweetsWithUserAuth(auth *models.TwitterAuthDetails, tweetID string, nextToken *string) (*structure.TweetRetweetsResponseWithUserAuth, error) {
	url := "https://api.twitter.com/1.1/statuses/retweeters/ids.json"
	paramsMap := map[string]string{
		"id":            tweetID,
		"stringify_ids": "true",
	}
	urlWithParams := fmt.Sprintf("%s?id=%s&stringify_ids=true", url, tweetID)
	if nextToken != nil {
		paramsMap["cursor"] = *nextToken
		urlWithParams = fmt.Sprintf("%s&cursor=%s", urlWithParams, *nextToken)
	}

	var response structure.TweetRetweetsResponseWithUserAuth
	err := executeTwtitterRequestWithUserOauth("GET", url, urlWithParams, auth.OAuthToken, auth.OAuthTokenSecret, nil, &response, paramsMap)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func FetchTwitterFollowersWithUserAuth(auth *models.TwitterAuthDetails, profileID string, nextToken *string) (*structure.TwitterFollowersResponseWithUserAuth, error) {
	url := "https://api.twitter.com/1.1/followers/ids.json"
	paramsMap := map[string]string{
		"user_id":       profileID,
		"stringify_ids": "true",
		"count":         "5000",
	}
	urlWithParams := fmt.Sprintf("%s?user_id=%s&stringify_ids=true&count=5000", url, profileID)
	if nextToken != nil {
		paramsMap["cursor"] = *nextToken
		urlWithParams = fmt.Sprintf("%s&cursor=%s", urlWithParams, *nextToken)
	}

	var response structure.TwitterFollowersResponseWithUserAuth
	err := executeTwtitterRequestWithUserOauth("GET", url, urlWithParams, auth.OAuthToken, auth.OAuthTokenSecret, nil, &response, paramsMap)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func executeTwitterRequest(method, endpoint string, requestData, destination interface{}) error {

	url := fmt.Sprintf("%s%s", twwitterAPIURL, endpoint)
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	var req *http.Request

	if requestData == nil {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(requestBody))
		if err != nil {
			return err
		}
	}

	req.Header.Set("Authorization", "Bearer "+utils.UseEnvOrDefault("TWITTER_BEARER_TOKEN", "sk-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
	req.Header.Set("Content-Type", "application/json")

	var response *http.Response
	response, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	responseCode := response.StatusCode
	if responseCode != 200 && responseCode != 201 {
		log.Print("error processing request: ", response)
		return errors.New("error processing request")
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(responseBody, destination)
}

func computeSignature(method, urlStr, baseString, oauthTokenSecret string) string {

	key := fmt.Sprintf("%s&%s", url.QueryEscape(os.Getenv("TWITTER_CONSUMER_SECRET")), url.QueryEscape(oauthTokenSecret))

	baseString = fmt.Sprintf("%s&%s&%s", method, url.QueryEscape(urlStr), url.QueryEscape(baseString))
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(baseString))

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func concatRequestParams(params map[string]string) string {
	var buffer bytes.Buffer
	for key, value := range params {
		buffer.WriteString(fmt.Sprintf("%s=%s&", key, value))
	}

	return strings.TrimSuffix(buffer.String(), "&")
}

func geneateOAuthHeader(params map[string]string) string {
	var buffer bytes.Buffer
	buffer.WriteString("OAuth ")

	for key, value := range params {
		buffer.WriteString(fmt.Sprintf("%s=\"%s\",", key, value))
	}

	return strings.TrimSuffix(buffer.String(), ",")
}

func mergeStringMaps(map1, map2 map[string]string) map[string]string {
	mergedMap := make(map[string]string)
	for k, v := range map1 {
		mergedMap[k] = v
	}

	for k, v := range map2 {
		mergedMap[k] = v
	}

	return mergedMap
}

func executeTwtitterRequestWithUserOauth(method, url, urlWithParams, oauthToken, oauthTokenSecret string, requestData, destination interface{}, paramsMap map[string]string) error {

	var params = map[string]string{}
	var req *http.Request
	var err error

	if requestData == nil {
		req, err = http.NewRequest(method, urlWithParams, nil)
		if err != nil {
			return err
		}
	} else {
		requestBody, err := json.Marshal(requestData)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(method, urlWithParams, bytes.NewBuffer(requestBody))
		if err != nil {
			return err
		}
	}

	params["oauth_consumer_key"] = os.Getenv("TWITTER_CONSUMER_KEY")
	params["oauth_nonce"] = utils.RandAlphaNumericRunes(32)
	params["oauth_signature_method"] = "HMAC-SHA1"
	params["oauth_timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["oauth_token"] = oauthToken
	params["oauth_version"] = "1.0"
	params = mergeStringMaps(params, paramsMap)
	params["oauth_signature"] = computeSignature(method, url, concatRequestParams(params), oauthTokenSecret)

	req.Header.Set("Authorization", geneateOAuthHeader(params))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var response *http.Response
	response, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	responseCode := response.StatusCode
	if responseCode != 200 && responseCode != 201 {
		log.Print("error processing request: ", response)
		return errors.New("error processing request")
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(responseBody, destination)
}

func FetchTweetDetails(link string) (*model.TweetDetails, error) {

	id, err := StripTweetIDFromLink(link)
	if err != nil {
		return nil, err
	}

	token := getGuestToken(link)
	if token == nil {
		return nil, errors.New("could not get guest token")
	}

	tweet, err := fetchTweetFromID(*id, *token)
	if err != nil {
		return nil, err
	}

	for _, tweet := range tweet.Data.ThreadedConversationWithInjectionsV2.Instructions[0].Entries {

		if strings.EqualFold(tweet.EntryID, fmt.Sprintf("tweet-%s", *id)) {
			return &model.TweetDetails{
				ProfilePhoto:  tweet.Content.ItemContent.TweetResults.Result.Core.UserResults.Result.Legacy.ProfileImageURLHTTPS,
				ProfileName:   tweet.Content.ItemContent.TweetResults.Result.Core.UserResults.Result.Legacy.Name,
				ProfileHandle: tweet.Content.ItemContent.TweetResults.Result.Core.UserResults.Result.Legacy.ScreenName,
				TweetText:     tweet.Content.ItemContent.TweetResults.Result.Legacy.FullText,
			}, nil
		}
	}

	return nil, errors.New("tweet not found")
}

func FetchUserDetails(userName string) (*model.UserDetails, error) {
	details, err := fetchDetailsFromUserName(userName)
	if err != nil {
		return nil, err
	}

	return &model.UserDetails{
		Name:     details.Data.Name,
		ID:       details.Data.ID,
		Username: details.Data.Username,
	}, nil
}

func FetchTweetByID(id string) (*structure.TweetResponse, error) {

	token := getGuestToken(twitterAuthEntryPoint)
	if token == nil {
		return nil, errors.New("could not get guest token")
	}

	tweet, err := fetchTweetFromID(id, *token)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func getGuestToken(url string) *string {

	body, err := fetchPrelimPage(url)
	if err != nil {
		log.Println(err)
		return nil
	}

	return utils.GetStrPtr(extractGuestToken(*body))
}

func extractGuestToken(body string) string {

	re := regexp.MustCompile(`document\.cookie="gt=(\d+)`)
	match := re.FindStringSubmatch(body)
	if len(match) > 1 {
		figures := match[1]
		return figures
	} else {
		return ""
	}

}

func StripTweetIDFromLink(link string) (*string, error) {

	idSegment := strings.Split(link, "/")[5]
	if len(idSegment) < 4 {
		return nil, errors.New("the tweet link provided seems to be invalid, please provide a valid link")
	}

	tweetID := strings.SplitN(idSegment, "?", 2)[0]
	return &tweetID, nil
}

func FetchTwitterAccessToken(token, verifier *string) (*models.TwitterAuthDetails, error) {

	params := url.Values{}
	params.Set("oauth_token", *token)
	params.Set("oauth_verifier", *verifier)
	url := fmt.Sprintf("https://api.twitter.com/oauth/access_token?%s", params.Encode())

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "PostmanRuntime/7.32.3")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "cc265e91-2d7d-42d2-841a-b6fc4cd7f43e")
	req.Header.Add("Host", "api.twitter.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "guest_id=v1%3A168728163881072847; guest_id_ads=v1%3A168728163881072847; guest_id_marketing=v1%3A168728163881072847; personalization_id=\"v1_Dpt96HQoIskLobYpTwUrQA==\"")
	req.Header.Add("Content-Length", "0")

	var res *http.Response
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	splitString := strings.Split(bytes.NewBuffer(body).String(), "&")
	var resp models.TwitterAuthDetails
	for _, s := range splitString {
		split := strings.Split(s, "=")
		if len(split) > 1 {
			switch split[0] {
			case "oauth_token":
				resp.OAuthToken = split[1]
			case "oauth_token_secret":
				resp.OAuthTokenSecret = split[1]
			case "user_id":
				resp.UserID = split[1]
			case "screen_name":
				resp.ScreenName = split[1]
			}
		}
	}

	return &resp, nil
}

func fetchTweetFromID(id, token string) (*structure.TweetResponse, error) {

	url := "https://twitter.com/i/api/graphql/tPRAv4UnqM9dOgDWggph7Q/TweetDetail?variables=%7B%22focalTweetId%22%3A%22" + id + "%22%2C%22with_rux_injections%22%3Afalse%2C%22includePromotedContent%22%3Atrue%2C%22withCommunity%22%3Atrue%2C%22withQuickPromoteEligibilityTweetFields%22%3Atrue%2C%22withBirdwatchNotes%22%3Afalse%2C%22withVoice%22%3Atrue%2C%22withV2Timeline%22%3Atrue%7D&features=%7B%22rweb_lists_timeline_redesign_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22tweetypie_unmention_optimization_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22graphql_is_translatable_rweb_tweet_is_translatable_enabled%22%3Atrue%2C%22view_counts_everywhere_api_enabled%22%3Atrue%2C%22longform_notetweets_consumption_enabled%22%3Atrue%2C%22tweet_awards_web_tipping_enabled%22%3Afalse%2C%22freedom_of_speech_not_reach_fetch_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Afalse%2C%22longform_notetweets_rich_text_read_enabled%22%3Atrue%2C%22longform_notetweets_inline_media_enabled%22%3Afalse%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D"

	var req *http.Request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Host", "twitter.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"")
	req.Header.Add("x-twitter-client-language", "en-GB")
	req.Header.Add("x-csrf-token", "81056247571840546c7c2f1874c5b61c")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-guest-token", token)
	req.Header.Add("x-twitter-active-user", "yes")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Referer", "https://twitter.com/tolusaba/status/1666176351022247937?s=46&t=GErzm5E5rICUqInKxbjCbA")
	req.Header.Add("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("Cookie", "guest_id=v1%3A168609151061125362; guest_id_ads=v1%3A168609151061125362; guest_id_marketing=v1%3A168609151061125362; personalization_id=\"v1_PjJQ9U+VY5UVOc4BAouPNQ==\"")

	var response *http.Response
	response, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseCode := response.StatusCode
	if responseCode != 200 && responseCode != 201 {
		log.Print("error processing request: ", response)
		return nil, errors.New("error processing request")
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp structure.TweetResponse
	err = json.Unmarshal(responseBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func fetchDetailsFromUserName(userName string) (*structure.UserDetailsResponse, error) {

	endpoint := fmt.Sprintf("users/by/username/%s", userName)

	var response structure.UserDetailsResponse
	err := executeTwitterRequest("GET", endpoint, nil, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil

}

func fetchPrelimPage(url string) (*string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	// Headers
	req.Header.Add("Host", "twitter.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("Cookie", "guest_id=v1%3A168609151061125362; guest_id_ads=v1%3A168609151061125362; guest_id_marketing=v1%3A168609151061125362; personalization_id=\"v1_PjJQ9U+VY5UVOc4BAouPNQ==\"")

	if err := req.ParseForm(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	// Fetch Request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	return utils.GetStrPtr(string(respBody)), nil
}
