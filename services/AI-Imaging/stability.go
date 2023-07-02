package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"inverse.so/graph/model"
	"inverse.so/structure"
	"inverse.so/utils"
)

const stabilityURL = "https://api.stability.ai/v1/generation/stable-diffusion-512-v2-1"

type StabilityService struct{}

func (s *StabilityService) GenerateImage(prompt string, style *model.AiImageStyle, number *int) ([]*model.ImageResponse, error) {

	resp, err := GenerateStabilityImage(prompt, style, number)
	if err != nil {
		return nil, err
	}

	var response []*model.ImageResponse
	for _, artifact := range resp.Artifacts {
		response = append(response, &model.ImageResponse{
			Image:  artifact.Base64,
			Format: model.ImageResolveFormaatBase64,
		})
	}

	return response, nil
}

func GenerateStabilityImage(prompt string, style *model.AiImageStyle, number *int) (*structure.StabilityImageResponse, error) {

	var stylePreset string = "fantasy-art"
	if style != nil {
		stylePreset = structure.ImageStyleMap[*style]
		prompt = saltPrompt(prompt, *style)
	}

	var n int = 1
	if number != nil {
		n = *number
	}

	requestData := &structure.StabilityImageRequest{
		Height: 512,
		Width:  512,
		TextPrompts: []structure.StabilityTextPrompt{
			{
				Text:   baseSalt,
				Weight: 1,
			},
			{
				Text:   prompt,
				Weight: 0.5,
			},
		},
		ClipGuidancePreset: "SIMPLE",
		Samples:            n,
		StylePreset:        stylePreset,
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
