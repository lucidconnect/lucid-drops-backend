package services

import "inverse.so/graph/model"

var promptSaltMap = map[model.AiImageStyle]string{
	model.AiImageStyleAnime:            "Concept art, pixiv-style, Sumi-e, portrait, ",
	model.AiImageStyleCinematic:        "Ultra realistic illustration, ",
	model.AiImageStyleDigitalArt:       "digital painting, ",
	model.AiImageStyleFantasyArt:       "artstation-style, postmodernism, ",
	model.AiImageStyleLineArt:          "Coloring Book style, line-art, ",
	model.AiImageStyleNeonPunk:         "Modernist, ",
	model.AiImageStyleOrigami:          "origami, ",
	model.AiImageStylePixelArt:         "pixel-art, portrait, ",
	model.AiImageStyleThreeDimensional: "hyperrealistic, unreal engine, portrait, ",
}

func saltPrompt(prompt string, style model.AiImageStyle) string {
	return promptSaltMap[style] + prompt
}

//resources for suture ref
//https://metaroids.com/lists/midjourney-art-styles-gigapack-free-200-prompt-keywords/
//https://stable-diffusion-art.com/prompts/
//https://stable-diffusion-art.com/how-to-come-up-with-good-prompts-for-ai-image-generation/#Some_good_keywords_for_you
//https://stable-diffusion-art.com/prompt-guide/
