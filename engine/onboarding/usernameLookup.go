package onboarding

import (
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
)

func CheckIfInverseNameIsAvailable(input *model.NewUsernameRegisgration) (bool, error) {
	_, err := engine.GetCreatorByInverseUsername(input.InverseUsername)
	if err == nil {
		return false, nil
	}

	return true, nil
}
