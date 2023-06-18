package onboarding

import (
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
)

func StoreUserAccountSignerAddress(input model.SignerInfo, authDetails *internal.AuthDetails) (bool, error) {

	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return false, err
	}

	userAccountSignerAddress := &models.SignerInfo{
		CreatorID:     creator.ID.String(),
		WalletAddress: input.Address,
		Signature:     input.Signature,
		Provider:      input.Provider,
	}

	err = engine.SaveModel(userAccountSignerAddress)
	if err != nil {
		return false, err
	}

	return true, nil
}
