package structure

import "github.com/lucidconnect/inverse/graph/model"

type MidJourneyImageRequest struct {
	Prompt      string `json:"prompt"`
	CallBackURL string `json:"callbackURL,omitempty"`
	Mode        string `json:"mode"`
}

type MidJourneyImageResponse struct {
	TaskID string `json:"taskid"`
}

type MidJourneyImageStatusRequest struct {
	TaskID   string `json:"taskId"`
	Position int    `json:"position,omitempty"`
}

type MidJourneyImageStatusResponse struct {
	Status     string `json:"status"`
	ImageURL   string `json:"imageURL"`
	Percentage int    `json:"percentage"`
}

func (m *MidJourneyImageStatusResponse) ToMidJourneyGraph() *model.ImageStatusResponse {
	return &model.ImageStatusResponse{
		Status:     &m.Status,
		Image:      &m.ImageURL,
		Percentage: &m.Percentage,
	}
}
