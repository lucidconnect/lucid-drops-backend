package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"inverse.so/structure"
	"inverse.so/utils"
)

const stabilityURL = "https://api.stability.ai/v1/generation/stable-diffusion-v1-5"

func GenerateStabilityImage(prompt, style string, number *int) (*structure.StabilityImageResponse, error) {

	var n int = 1
	if number != nil {
		n = *number
	}

	requestData := &structure.StabilityImageRequest{
		Height: 512,
		Width:  512,
		TextPrompts: []structure.StabilityTextPrompt{
			{
				Text:   prompt,
				Weight: 1,
			},
		},
		Samples:     n,
		StylePreset: style,
	}

	var response structure.StabilityImageResponse
	err := executeStabilityRequest("POST", "text-to-image", requestData, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Artifacts) == 0 {
		return nil, errors.New("no data returned")
	}

	return &response, nil
}

func executeStabilityRequest(method, endpoint string, requestData, destination interface{}) error {

	reqUrl := fmt.Sprintf("%s/%s", stabilityURL, endpoint)

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	var req *http.Request

	if requestData == nil {
		req, err = http.NewRequest(method, reqUrl, nil)
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest(method, reqUrl, bytes.NewBuffer(requestBody))
		if err != nil {
			return err
		}
	}

	req.Header.Set("Authorization", utils.UseEnvOrDefault("STABILITY_API_KEY", "sk-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Organization", utils.UseEnvOrDefault("STABILITY_ORG_ID", "sk-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

	var response *http.Response
	log.Print("request: ", req)
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
