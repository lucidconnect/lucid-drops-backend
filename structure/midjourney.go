package structure

type MidJourneyImageRequest struct {
	Prompt      string `json:"prompt"`
	CallBackURL string `json:"callbackURL,omitempty"`
	Mode        string `json:"mode"`
}

type MidJourneyImageResponse struct {
	TaskID string `json:"taskid"`
}
