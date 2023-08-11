package onboarding

import (
	"errors"
	"fmt"

	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/models"
	"inverse.so/utils"
)

func RegisterInverseUsername(address string, input *model.NewUsernameRegisgration) (*model.CreatorDetails, error) {
	_, err := engine.GetCreatorByInverseUsername(input.InverseUsername)
	if err == nil {
		return nil, errors.New("inverse name isn't available")
	}

	cachedCreator, storedCreatorErr := engine.GetCreatorByAddress(address)
	if storedCreatorErr != nil {
		newCreator := models.Creator{WalletAddress: address, InverseUsername: utils.GetStringPtr(input.InverseUsername)}
		creationErr := dbutils.DB.Create(&newCreator).Error
		if creationErr != nil {
			return nil, creationErr
		}
		return newCreator.ToGraphData(), nil
	}

	if cachedCreator.InverseUsername != nil {
		return nil, fmt.Errorf("creator already has (%s) as there inverse name", *cachedCreator.InverseUsername)
	}

	cachedCreator.InverseUsername = utils.GetStringPtr(input.InverseUsername)
	creationErr := dbutils.DB.Save(&cachedCreator).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return cachedCreator.ToGraphData(), nil
}
