package aiimages

import (
	"inverse.so/graph/model"
)

// var numberOfImagesToGenerate int = 6

func GetImageSuggestions(prompt string, presets *model.AiImageStyle) ([]string, error) {
	return []string{
		"https://randnft.vercel.app/api/rand",
		"https://randnft.vercel.app/api/rand",
		"https://randnft.vercel.app/api/rand",
		"https://randnft.vercel.app/api/rand",
		"https://randnft.vercel.app/api/rand",
		"https://randnft.vercel.app/api/rand",
	}, nil

	// var stylePrompt string = "anime"

	// if presets != nil {
	// 	stylePrompt = string(*presets)
	// }

	// stabilityResponse, err := services.GenerateStabilityImage(prompt, stylePrompt, &numberOfImagesToGenerate)
	// if err != nil {
	// 	return nil, err
	// }

	// aiImages := make([]string, len(stabilityResponse.Artifacts))

	// for idx, image := range stabilityResponse.Artifacts {
	// 	aiImages[idx] = image.Base64
	// }

	// return aiImages, nil
}
