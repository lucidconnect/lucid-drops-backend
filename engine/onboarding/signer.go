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

	noSignature := "NONE" // TODO use aa wallet signatures to authorize third party signer
	if input.Signature == nil {
		input.Signature = &noSignature
	}

	altSigner, err := engine.GetAltSignerByCreatorID(creator.ID.String())
	if err != nil {
		altSigner = &models.SignerInfo{
			CreatorID:     creator.ID.String(),
			WalletAddress: input.Address,
			Signature:     input.Signature,
			Provider:      input.Provider,
		}
	} else {
		altSigner.WalletAddress = input.Address
		altSigner.Provider = input.Provider
		altSigner.Signature = input.Signature
	}

	alterr := engine.SaveModel(altSigner)
	if alterr != nil {
		return false, err
	}

	return true, nil
}
