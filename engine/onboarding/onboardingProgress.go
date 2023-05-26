package onboarding

import (
	"inverse.so/engine"
	"inverse.so/graph/model"
)

func GetOnboardinProgress(address string) (*model.OnboardingProgress, error) {
	cachedCreator, err := engine.GetCreatorByAddress(address)
	if err != nil {
		return &model.OnboardingProgress{
			RegisterdInverseUsername: false,
		}, nil
	}

	return &model.OnboardingProgress{
		RegisterdInverseUsername: cachedCreator.InverseUsername != nil,
	}, nil

}
