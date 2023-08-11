package aiimages

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

const openAIURL = "https://api.openai.com/v1"

type OpenAIService struct{}

func (o *OpenAIService) GenerateImage(prompt string, style *model.AiImageStyle, number *int) ([]*model.ImageResponse, error) {

	resp, err := GenerateOpenAIImage(prompt, style, number)
	if err != nil {
		return nil, err
	}

	var response []*model.ImageResponse
	for _, image := range resp.Data {
		response = append(response, &model.ImageResponse{
			Image:  image.URL,
			Format: model.ImageResolveFormaatURL,
		})
	}

	return response, nil
}

func GenerateOpenAIImage(prompt string, style *model.AiImageStyle, number *int) (*structure.ImageResponse, error) {

	var n int = 1
	if number != nil {
		n = *number
	}

	if style != nil {
		prompt = saltPrompt(prompt, *style)
	}

	requestData := &structure.ImageRequest{
		Prompt: prompt,
		N:      n,
		Size:   "1024x1024",
	}

	var response structure.ImageResponse
	err := executeOpenAIRequest("POST", "images/generations", requestData, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Data) == 0 {
		return nil, errors.New("no data returned")
	}

	return &response, nil
}

func executeOpenAIRequest(method, endpoint string, requestData, destination interface{}) error {

	url := fmt.Sprintf("%s/%s", openAIURL, endpoint)
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

	req.Header.Set("Authorization", "Bearer "+utils.UseEnvOrDefault("OPENAI_API_KEY", "sk-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
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
