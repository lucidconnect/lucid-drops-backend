package aiimages

import (
	"github.com/lucidconnect/inverse/graph/model"
)

var numberOfImagesToGenerate int = 4

type LLMProvider interface {
	GenerateImage(prompt string, style *model.AiImageStyle, numberOfImagesToGenerate *int) ([]*model.ImageResponse, error)
}

func GetImageSuggestions(prompt string, presets *model.AiImageStyle) ([]*model.ImageResponse, error) {

	provider := MidJourneyService{}
	return provider.GenerateImage(prompt, presets, &numberOfImagesToGenerate)

}
