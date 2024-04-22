package neynar

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/lucidconnect/inverse/drops"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/utils"
	"github.com/rs/zerolog/log"
)

/**
* User
	- use fid for validations
	- fetch a creators following
	- validate if a user is following a given farcaster account

* Casts
	- fetch likes for a given cast
	- fetch replies to a given cast
	- fetch quotes to a given cast

*Channels
	- fetch followers of a given channel
	- validate if a user follows a given channel
*/

type NeynarClient struct {
	client       *http.Client
	apiKey       string
	neynarUrl    string
	farcasterHub string
	// errMsg       error
}

type Option func(*NeynarClient)

func NewNeynarClient(options ...Option) (*NeynarClient, error) {
	neynarClient := &NeynarClient{
		client:       &http.Client{},
		apiKey:       "",
		neynarUrl:    "https://api.neynar.com",
		farcasterHub: "https://hub-api.neynar.com",
	}

	for _, opt := range options {
		opt(neynarClient)
	}

	if neynarClient.apiKey == "" {
		return nil, errors.New("trying to initialise neynar without an api key")
	}

	return neynarClient, nil
}

// WithTimeout is a functional option to set the HTTP client timeout.
func WithNeynarApiKey(key string) Option {
	return func(c *NeynarClient) {
		c.apiKey = key
	}
}

func (nc *NeynarClient) ValidateFarcasterCriteriaForDrop(farcasterAddress string, drop drops.Drop) (*model.ValidationRespoonse, error) {
	resp := &model.ValidationRespoonse{
		Valid: false,
	}

	criteria := drop.FarcasterCriteria
	requiredCriteriaTypes := strings.Split(criteria.CriteriaType, ",")
	userFid, err := nc.FetchFarcasterUserFidByEthAddress(farcasterAddress)
	if err != nil {
		return resp, err
	}

	for _, criteriaType := range requiredCriteriaTypes {
		if criteriaType == model.ClaimCriteriaTypeFarcasterInteractions.String() {
			for _, interaction := range drops.InteractionsToArr(criteria.Interactions) {
				switch *interaction {
				case model.InteractionTypeReplies:
					if !nc.validateFarcasterReplyCriteria(userFid, *criteria) {
						resp.Message = utils.GetStrPtr("farcaster account does not meet the reply criteria")
						return resp, errors.New("farcaster account does not meet the reply criteria")
					}
				case model.InteractionTypeRecasts:
					if !nc.validateFarcasterRecastCriteria(userFid, *criteria) {
						resp.Message = utils.GetStrPtr("farcaster account does not meet the recast criteria")
						return resp, errors.New("farcaster account does not meet the recast criteria")
					}
				case model.InteractionTypeLikes:
					if !nc.validateFarcasterLikeCriteria(userFid, *criteria) {
						resp.Message = utils.GetStrPtr("farcaster account does not meet the like criteria")
						return resp, errors.New("farcaster account does not meet the `like` criteria")
					}
				}
			}
		}

		if criteriaType == model.ClaimCriteriaTypeFarcasterFollowing.String() {
			if !nc.validateFarcasterAccountFollowerCriteria(userFid, *criteria) {
				resp.Message = utils.GetStrPtr("farcaster account does not meet the follower criteria")
				return resp, errors.New("farcaster account does not meet the follower criteria")
			}
		}

		if criteriaType == model.ClaimCriteriaTypeFarcasterChannel.String() {
			if !nc.validateFarcasterChannelFollowerCriteria(userFid, *criteria) {
				resp.Message = utils.GetStrPtr("farcaster account must be following the required channel")
				return resp, errors.New("farcaster account does not meet the channel follower criteria")
			}
		}
	}
	resp.Valid = true
	return resp, nil
}

// Returns a list of relevant followers for a given fid
func (nc *NeynarClient) FetchRelvantFollowers(fid int32) ([]RelevantFollowersDehydrated, error) {
	url, err := url.Parse(fmt.Sprintf("%v/v2/farcaster/followers/relevant", nc.neynarUrl))
	if err != nil {
		return nil, err
	}
	query := url.Query()
	query.Add("target_fid", fmt.Sprint(fid))
	query.Add("viewer_fid", fmt.Sprint(fid))

	url.RawQuery = query.Encode()
	response, err := nc.makeRequest(http.MethodGet, url.String(), "", nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	followers, err := decodeRelevantFollowers(response.Body)
	if err != nil {
		return nil, err
	}

	return followers.AllRelevantFollowersDehydrated, nil
}

func (nc *NeynarClient) RetrieveCastByUrl(castUrl string) (Cast, error) {
	url, err := url.Parse(fmt.Sprintf("%v/v2/farcaster/cast", nc.neynarUrl))
	if err != nil {
		return Cast{}, err
	}
	query := url.Query()
	query.Add("identifier", castUrl)
	query.Add("type", "url")

	url.RawQuery = query.Encode()
	response, err := nc.makeRequest(http.MethodGet, url.String(), "", nil)
	if err != nil {
		return Cast{}, err
	}
	defer response.Body.Close()

	cast, err := decodeCastObject(response.Body)
	if err != nil {
		return Cast{}, err
	}
	return cast, nil
}

// Now deprecated, use v2/cast/conversation. Gets all casts, including root cast and all replies for a given thread hash. No limit the depth of replies.
func (nc *NeynarClient) RetrieveCastsByThreadHash(hash string) ([]Cast, error) {
	url, err := url.Parse(fmt.Sprintf("%v/v1/farcaster/all-casts-in-thread", nc.neynarUrl))
	if err != nil {
		return nil, err
	}
	query := url.Query()
	query.Add("threadHash", hash)

	url.RawQuery = query.Encode()
	response, err := nc.makeRequest(http.MethodGet, url.String(), "", nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	cast, err := decodeThreadCasts(response.Body)
	if err != nil {
		return nil, err
	}

	return cast, nil
}

func (nc *NeynarClient) RetrieveChannelFollowers(channelID string, fid int32, cursor string) (ChannelFollowers, error) {
	url, err := url.Parse(fmt.Sprintf("%v/v2/farcaster/channel/followers", nc.neynarUrl))
	if err != nil {
		return ChannelFollowers{}, err
	}
	query := url.Query()
	query.Add("id", channelID)
	query.Add("limit", "1000")

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	url.RawQuery = query.Encode()
	response, err := nc.makeRequest(http.MethodGet, url.String(), "", nil)
	if err != nil {
		return ChannelFollowers{}, err
	}
	defer response.Body.Close()

	followers, err := decodeChannelFollowers(response.Body)
	if err != nil {
		return ChannelFollowers{}, err
	}

	return followers, nil
}

func (nc *NeynarClient) FetchFarcasterUserFidByEthAddress(address string) (int32, error) {
	url, err := url.Parse(fmt.Sprintf("%v/v2/farcaster/user/bulk-by-address", nc.neynarUrl))
	if err != nil {
		return 0, err
	}
	query := url.Query()
	query.Add("addresses", address)

	url.RawQuery = query.Encode()
	response, err := nc.makeRequest(http.MethodGet, url.String(), "", nil)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	user, err := decodeFarcasterUser(response.Body, address)
	if err != nil {
		return 0, err
	}

	return user.Fid, err
}

func (nc *NeynarClient) FetchFarcasterUserByUsername(username string) (int32, error) {
	url, err := url.Parse(fmt.Sprintf("%v/v1/farcaster/user-by-username", nc.neynarUrl))
	if err != nil {
		return 0, err
	}

	query := url.Query()
	query.Add("username", username)
	url.RawQuery = query.Encode()
	response, err := nc.makeRequest(http.MethodGet, url.String(), "", nil)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	responseObj := map[string]map[string]map[string]any{}
	if err = json.NewDecoder(response.Body).Decode(&responseObj); err != nil {
		return 0, err
	}

	fid := (responseObj["result"]["user"]["fid"]).(float64)

	return int32(fid), nil
}

func (nc *NeynarClient) retrieveUsersChannels(fid int32, cursor string) (UserChannels, error) {
	url, err := url.Parse(fmt.Sprintf("%v/v2/farcaster/user/channels", nc.neynarUrl))
	if err != nil {
		return UserChannels{}, err
	}
	query := url.Query()
	query.Add("fid", strconv.FormatInt(int64(fid), 10))
	query.Add("limit", "100")
	if cursor != "" {
		query.Add("cursor", cursor)
	}

	url.RawQuery = query.Encode()
	fmt.Println(url.String())

	response, err := nc.makeRequest(http.MethodGet, url.String(), "", nil)
	if err != nil {
		return UserChannels{}, err
	}
	defer response.Body.Close()

	channels, err := decodeUserChannels(response.Body)
	if err != nil {
		return UserChannels{}, err
	}

	return channels, err
}

func (nc *NeynarClient) makeRequest(method, url, contentType string, body io.Reader) (*http.Response, error) {
	httpRequest, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("api_key", nc.apiKey)
	httpRequest.Header.Set("Content-Type", contentType)
	httpRequest.Header.Set("accept", "appication/json")

	resp, err := nc.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var errResponse ErrorResponse
	if resp.StatusCode != 200 {
		if err = json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			return nil, err
		}
		err = errors.New(errResponse.Message)
		return nil, err
	}

	return resp, nil
}

func (nc *NeynarClient) validateFarcasterChannelFollowerCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	var userChannels []string
	channels := strings.Split(criteria.ChannelID, ",")

	channelObject, err := nc.retrieveUsersChannels(fid, "")
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	userChannels = appendUserChannels(channelObject)

	for {
		if channelObject.Next.Cursor == "" {
			break
		} else {
			channelObject, err := nc.retrieveUsersChannels(fid, "")
			if err != nil {
				log.Err(err).Caller().Send()
				break
			}
			userChannels = append(userChannels, appendUserChannels(channelObject)...)
		}
	}
	sort.Strings(userChannels)
	fmt.Printf("user %v channels: %v", fid, userChannels)

	for _, channel := range channels {
		idx := sort.SearchStrings(userChannels, channel)
		if userChannels[idx] == channel {
			return true
		}
	}
	return false
}

func (nc *NeynarClient) validateFarcasterLikeCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	cast, err := nc.RetrieveCastByUrl(criteria.CastUrl)
	if err != nil {
		return false
	}
	fmt.Println("Likes - ", cast)
	for _, like := range cast.Reactions.Likes {
		if like.Fid == fid {
			return true
		}
	}

	return false
}

func (nc *NeynarClient) validateFarcasterRecastCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	cast, err := nc.RetrieveCastByUrl(criteria.CastUrl)
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

func (nc *NeynarClient) validateFarcasterReplyCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	rootCast, err := nc.RetrieveCastByUrl(criteria.CastUrl)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	casts, err := nc.RetrieveCastsByThreadHash(rootCast.Hash)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}

	for _, cast := range casts {
		if cast.Author.Fid == fid {
			return true
		}
	}
	return false
}

func (nc *NeynarClient) validateFarcasterAccountFollowerCriteria(fid int32, criteria drops.FarcasterCriteria) bool {
	var followers []RelevantFollowersDehydrated

	creatorFid, err := strconv.Atoi(criteria.FarcasterProfileID)
	if err != nil {
		log.Err(err).Caller().Send()
		return false
	}
	followers, err = nc.FetchRelvantFollowers(int32(creatorFid))
	if err != nil {
		return false
	}

	fmt.Println(followers)

	for _, follower := range followers {
		if follower.User.Fid == fid {
			return true
		}
	}
	return false
}

func decodeRelevantFollowers(response io.ReadCloser) (FarcasterFollowers, error) {
	var err error
	relevantFollowers := FarcasterFollowers{}

	if err = json.NewDecoder(response).Decode(&relevantFollowers); err != nil {
		err = fmt.Errorf("failed to decode response body: %v", err)
		return FarcasterFollowers{}, err
	}

	return relevantFollowers, nil
}

func decodeCastObject(response io.ReadCloser) (Cast, error) {
	var err error
	responseBody := RetrieveCastResponse{}

	if err = json.NewDecoder(response).Decode(&responseBody); err != nil {
		err = fmt.Errorf("failed to decode response body: %v", err)
		return Cast{}, err
	}
	return responseBody.Cast, nil
}

func decodeChannelFollowers(response io.ReadCloser) (ChannelFollowers, error) {
	var err error
	followers := ChannelFollowers{}

	if err = json.NewDecoder(response).Decode(&followers); err != nil {
		err = fmt.Errorf("failed to decode response body: %v", err)
		return ChannelFollowers{}, err
	}
	return followers, nil
}

func decodeUserChannels(response io.ReadCloser) (UserChannels, error) {
	var err error
	channels := UserChannels{}

	if err = json.NewDecoder(response).Decode(&channels); err != nil {
		err = fmt.Errorf("failed to decode response body: %v", err)
		return UserChannels{}, err
	}

	return channels, nil
}

func decodeThreadCasts(response io.ReadCloser) ([]Cast, error) {
	var err error
	cast := ThreadCasts{}

	if err = json.NewDecoder(response).Decode(&cast); err != nil {
		err = fmt.Errorf("failed to decode response body: %v", err)
		return nil, err
	}

	return cast.Result.Casts, nil
}

func decodeFarcasterUser(response io.ReadCloser, address string) (UserDehydrated, error) {
	var err error

	responseBody := map[string][]UserDehydrated{}
	// u := &UserDehydrated{}
	if err = json.NewDecoder(response).Decode(&responseBody); err != nil {
		err = fmt.Errorf("failed to decode response body: %v", err)
		return UserDehydrated{}, err
	}
	fmt.Println(responseBody)
	if len(responseBody) == 0 {
		return UserDehydrated{}, fmt.Errorf("address %v is not connected with a farcaster account", address)
	}
	userI := responseBody[strings.ToLower(address)][0]
	return userI, nil
}

func appendUserChannels(userChannels UserChannels) []string {
	var channelIDs []string

	for _, channel := range userChannels.Channels {
		channelIDs = append(channelIDs, channel.Id)
	}

	return channelIDs
}
