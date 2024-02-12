package onboarding

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/engine/mobile"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/utils"
)

func RegisterInverseUsername(address common.Address, input *model.NewUsernameRegisgration) (*model.CreatorDetails, error) {
	_, err := engine.GetCreatorByInverseUsername(input.InverseUsername)
	if err == nil {
		return nil, errors.New("inverse name isn't available")
	}
	cachedCreator, storedCreatorErr := engine.GetCreatorByAddress(address)
	if storedCreatorErr != nil {
		newCreator := models.Creator{WalletAddress: address.String(), InverseUsername: utils.GetStringPtr(input.InverseUsername)}
		creationErr := dbutils.DB.Create(&newCreator).Error
		if creationErr != nil {
			return nil, creationErr
		}

		accountDetails, generationErr := mobile.GenerateRandomEthAddress()
		if generationErr != nil {
			return nil, generationErr
		}

		aaWallet := common.HexToAddress(input.AaWallet)
		newAltSigner := models.SignerInfo{
			CreatorID:     newCreator.ID.String(),
			WalletAddress: aaWallet.String(),
			Provider:      model.SignerProviderConnectKit,
			AltPublicKey:  accountDetails.PublicKey,
			AltPrivateKey: accountDetails.PrivateKey,
		}

		altSignerErr := dbutils.DB.Create(&newAltSigner).Error
		if altSignerErr != nil {
			return nil, altSignerErr
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

	altSigner, err := engine.GetAltSignerByCreatorID(cachedCreator.ID.String())
	if err != nil {
		aaWallet := common.HexToAddress(input.AaWallet)

		altSigner = &models.SignerInfo{
			CreatorID:     cachedCreator.ID.String(),
			WalletAddress: aaWallet.String(),
			Provider:      model.SignerProviderConnectKit,
		}
	} else {
		aaWallet := common.HexToAddress(input.AaWallet)
		altSigner.WalletAddress = aaWallet.String()
		altSigner.Provider = model.SignerProviderConnectKit
	}

	err = engine.SaveModel(altSigner)
	if err != nil {
		return nil, err
	}

	return cachedCreator.ToGraphData(), nil
}
