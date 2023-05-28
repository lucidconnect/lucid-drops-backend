package structure

import "inverse.so/graph/model"

type StabilityImageRequest struct {
	Height      int                   `json:"height"`
	Width       int                   `json:"width"`
	TextPrompts []StabilityTextPrompt `json:"text_prompts"`
	Samples     int                   `json:"samples"`
	StylePreset string                `json:"style_preset,omitempty"`
}

type StabilityImageResponse struct {
	Artifacts []StabilityResponse `json:"artifacts"`
}

type StabilityResponse struct {
	Base64       string  `json:"base64"`
	FinishReason string  `json:"finishReason"`
	Seed         float64 `json:"seed"`
}

type StabilityTextPrompt struct {
	Text   string  `json:"text"`
	Weight float64 `json:"weight"`
}

var ImageStyleMap = map[model.AiImageStyle]string{
	model.AiImageStyleAnime:            "anime",
	model.AiImageStyleCinematic:        "cinematic",
	model.AiImageStyleDigitalArt:       "digital-art",
	model.AiImageStyleFantasyArt:       "fantasy-art",
	model.AiImageStyleLineArt:          "line-art",
	model.AiImageStyleNeonPunk:         "neon-punk",
	model.AiImageStyleOrigami:          "origami",
	model.AiImageStylePixelArt:         "pixel-art",
	model.AiImageStyleThreeDimensional: "3d-model",
}
