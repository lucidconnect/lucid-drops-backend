package auth

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/engine/onboarding"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/magic"
	"github.com/lucidconnect/inverse/models"
)

const (
	InverseAuthSignatureMessage = "Authorize Inverse Authentication"
)

func CreateJWTToken(input *model.CreateJWTTokenInput) (*model.JWTCreationResponse, error) {
	message := "\x19Ethereum Signed Message:\n32" + string(InverseAuthSignatureMessage)
	signer, err := magic.GetMeTheSignerOfThisMessage(message, input.Signature)
	if err != nil {
		return nil, err
	}

	castedAddress := common.HexToAddress(input.Address)
	if *signer != castedAddress {
		return nil, errors.New("the signature provided is invalid")
	}

	jwt, err := internal.GenerateJWT(input.Address)
	if err != nil {
		return nil, errors.New("jwt creation has failed")
	}

	parsedAddress := common.HexToAddress(input.Address)
	creatorInfo, err := onboarding.CreateCreatorProfileIfAddressIsMissing(parsedAddress)
	if err != nil {
		return nil, err
	}

	_, err = engine.GetAltSignerByCreatorID(creatorInfo.ID.String())
	if err != nil {
		aaWallet := common.HexToAddress(input.AaWallet)
		altSigner := &models.SignerInfo{
			CreatorID:     creatorInfo.ID.String(),
			WalletAddress: aaWallet.String(),
			Provider:      model.SignerProviderConnectKit,
		}

		err = engine.SaveModel(altSigner)
		if err != nil {
			return &model.JWTCreationResponse{
				Token: jwt,
			}, nil
		}
	}

	return &model.JWTCreationResponse{
		Token: jwt,
	}, nil
}
