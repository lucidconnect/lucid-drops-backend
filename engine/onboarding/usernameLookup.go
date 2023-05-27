package onboarding

import (
	"inverse.so/engine"
	"inverse.so/graph/model"
)

func CheckIfInverseNameIsAvailable(input *model.NewUsernameRegisgration) (bool, error) {
	_, err := engine.GetCreatorByInverseUsername(input.InverseUsername)
	if err != nil {
		return false, nil
	}

	return true, nil
}
