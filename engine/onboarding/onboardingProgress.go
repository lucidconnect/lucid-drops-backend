package onboarding

import (
	"inverse.so/engine"
	"inverse.so/graph/model"
)

func GetOnboardinProgress(address string) (*model.OnboardingProgress, error) {
	cachedCreator, err := engine.GetCreatorByAddress(address)
	if err != nil {
		return &model.OnboardingProgress{
			Creator:                  nil,
			RegisterdInverseUsername: false,
		}, nil
	}

	return &model.OnboardingProgress{
		Creator:                  cachedCreator.ToGraphData(),
		RegisterdInverseUsername: cachedCreator.InverseUsername != nil,
	}, nil

}
