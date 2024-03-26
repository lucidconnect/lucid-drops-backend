package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func createMintUrl(item, imagUrl, contract string) (string, error) {
	baseurl := os.Getenv("FRAME_SERVER")
	url := fmt.Sprintf("%v/createframe", baseurl)

	type createFrameRequest struct {
		DropId     string `json:"dropId"`
		ImageUrl   string `json:"imageUrl"`
		Collection string `json:"collection"`
	}

	request := createFrameRequest{
		DropId:     item,
		ImageUrl:   imagUrl,
		Collection: contract,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	httpRequest, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	httpRequest.Header.Set("Content-Type", "appication/json")

	res, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var frameUrl string
	if err = json.NewDecoder(res.Body).Decode(&frameUrl); err != nil {
		return "", err
	}

	return frameUrl, nil
}
