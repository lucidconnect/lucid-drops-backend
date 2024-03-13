package neynar

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/** TODO
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

// Returns a list of relevant followers for a given fid
func (nc *NeynarClient) FetchRelvantFollowers(fid uint32) ([]RelevantFollowersDehydrated, error) {
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

func (nc *NeynarClient) RetrieveCastByThreadHash(hash string) (Cast, error) {
	url, err := url.Parse(fmt.Sprintf("%v/v1/farcaster/all-casts-in-thread", nc.neynarUrl))
	if err != nil {
		return Cast{}, err
	}
	query := url.Query()
	query.Add("threadHash", hash)

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

	return resp, nil
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
	cast := Cast{}

	if err = json.NewDecoder(response).Decode(&cast); err != nil {
		err = fmt.Errorf("failed to decode response body: %v", err)
		return Cast{}, err
	}
	return cast, nil
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
