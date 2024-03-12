package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type AlchemyClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

type Option func(*AlchemyClient)

type AlchemyGetOwnersForNftResponse struct {
	Owners  []string `json:"owners"`
	PageKey string   `json:"pageKey"`
}

func NewAlchemyClient(options ...Option) (*AlchemyClient, error) {
	alchemyClient := &AlchemyClient{
		client: &http.Client{},
	}

	for _, opt := range options {
		opt(alchemyClient)
	}

	if alchemyClient.apiKey == "" {
		return nil, errors.New("trying to initialise alcheny without an api key")
	}

	return alchemyClient, nil
}

func WithApiKey(key string) Option {
	return func(ac *AlchemyClient) {
		ac.apiKey = key
	}
}

func WithUrl(url string) Option {
	return func(ac *AlchemyClient) {
		ac.baseUrl = url
	}
}

func (ac *AlchemyClient) GetOwnersForNft(contractAddress, tokenId string) ([]string, error) {
	urlString := fmt.Sprintf("%v/nft/v3/%v/getOwnersForNFT", ac.baseUrl, ac.apiKey)

	reqUrl, err := url.Parse(urlString)
	if err != nil {
		err = fmt.Errorf("GetOwnersForNft(): %v", err)
		return nil, err
	}

	query := reqUrl.Query()
	query.Add("contractAddress", contractAddress)
	query.Add("tokenId", tokenId)

	reqUrl.RawQuery = query.Encode()
	fmt.Println(reqUrl.String())
	httpRequest, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		err = fmt.Errorf("GetOwnersForNft(): %v", err)
		return nil, err
	}
	res, err := ac.client.Do(httpRequest)
	if err != nil {
		err = fmt.Errorf("GetOwnersForNft(): %v", err)
		return nil, err
	}
	defer res.Body.Close()

	var getOwnersForNFTResponse *AlchemyGetOwnersForNftResponse
	if err = json.NewDecoder(res.Body).Decode(&getOwnersForNFTResponse); err != nil {
		err = fmt.Errorf("GetOwnersForNft() - json decoding: %v", err)
		return nil, err
	}

	if getOwnersForNFTResponse.PageKey != "" {
		// implement pagination
	}

	return getOwnersForNFTResponse.Owners, nil
}
