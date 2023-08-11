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

const midJourneyURLURL = "https://api.midjourneyapi.io/v2"

type MidJourneyService struct{}

func (m *MidJourneyService) GenerateImage(prompt string, style *model.AiImageStyle, number *int) ([]*model.ImageResponse, error) {

	resp, err := GenerateMidJourneyImage(prompt, style, number)
	if err != nil {
		return nil, err
	}

	var response []*model.ImageResponse
	response = append(response, &model.ImageResponse{
		Image:  "",
		Format: model.ImageResolveFormaatURL,
		TaskID: utils.GetStrPtr(resp.TaskID),
	})

	return response, nil
}

func GenerateMidJourneyImage(prompt string, style *model.AiImageStyle, number *int) (*structure.MidJourneyImageResponse, error) {

	// var n int = 1
	// if number != nil {
	// 	n = *number
	// }

	if style != nil {
		prompt = saltPrompt(prompt, *style)
	}

	requestData := &structure.MidJourneyImageRequest{
		Prompt:      prompt,
		CallBackURL: "",
		Mode:        "fast",
	}

	var response structure.MidJourneyImageResponse
	err := executeMidjourneyRequest("POST", "/imagine", requestData, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func QueryMidJourneyTaskID(taskID string, position *int) (*model.ImageStatusResponse, error) {

	requestData := &structure.MidJourneyImageStatusRequest{
		TaskID: taskID,
	}

	endpoint := "/result"
	if position != nil {
		requestData.Position = *position
		endpoint = "/upscale"
	}

	var response structure.MidJourneyImageStatusResponse
	err := executeMidjourneyRequest("POST", endpoint, requestData, &response)
	if err != nil {
		return nil, err
	}

	return response.ToMidJourneyGraph(), nil
}

func executeMidjourneyRequest(method, endpoint string, requestData, destination interface{}) error {

	url := fmt.Sprintf("%s%s", midJourneyURLURL, endpoint)
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

	req.Header.Set("Authorization", utils.UseEnvOrDefault("MIDJOURNEY_API_KEY", "sk-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
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
