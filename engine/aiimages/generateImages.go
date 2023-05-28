package aiimages

import (
	"inverse.so/graph/model"
	services "inverse.so/services/AI-Imaging"
)

var numberOfImagesToGenerate int = 6

type LLMProvider interface {
	GenerateImage(prompt string, style *model.AiImageStyle, numberOfImagesToGenerate *int) ([]*model.ImageResponse, error)
}

func GetImageSuggestions(prompt string, presets *model.AiImageStyle) ([]*model.ImageResponse, error) {
	
	provider := services.StabilityService{}
	return provider.GenerateImage(prompt, presets, &numberOfImagesToGenerate)

}
