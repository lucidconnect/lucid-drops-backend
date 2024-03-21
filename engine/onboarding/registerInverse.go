package onboarding

import (
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/utils"
	"gorm.io/gorm"
)

func RegisterInverseUsername(address common.Address, input *model.NewUsernameRegisgration) (*model.CreatorDetails, error) {
	_, err := engine.GetCreatorByInverseUsername(input.InverseUsername)
	if err == nil {
		return nil, errors.New("inverse name isn't available")
	}

	// move this flow to register username
	creatorInfo, err := CreateCreatorProfile(address)
	if err != nil {
		log.Println("creator profile: ", err)
		return nil, err
	}

	// if cachedCreator.InverseUsername != nil {
	// 	return nil, fmt.Errorf("creator already has (%s) as there inverse name", *cachedCreator.InverseUsername)
	// }

	creatorInfo.InverseUsername = utils.GetStringPtr(input.InverseUsername)
	creatorInfo.AAWalletAddress = input.AaWallet
	if input.ExternalWalletAddress != nil {
		creatorInfo.ExternalWalletAddress = *input.ExternalWalletAddress
	}

	creationErr := dbutils.DB.Save(&creatorInfo).Error
	if creationErr != nil {
		return nil, creationErr
	}

	// altSigner, err := engine.GetAltSignerByCreatorID(creatorInfo.ID.String())
	// if err != nil {
	// aaWallet := common.HexToAddress(input.AaWallet)

	altSigner := &models.SignerInfo{
		CreatorID:     creatorInfo.ID.String(),
		WalletAddress: input.AaWallet,
		Provider:      model.SignerProviderConnectKit,
	}
	// } else {
	// 	aaWallet := common.HexToAddress(input.AaWallet)
	// 	altSigner.WalletAddress = aaWallet.String()
	// 	altSigner.Provider = model.SignerProviderConnectKit
	// }

	err = engine.SaveModel(altSigner)
	if err != nil {
		return nil, err
	}

	return creatorInfo.ToGraphData(), nil
}

func CreateCreatorProfile(address common.Address) (*models.Creator, error) {
	cachedCreator, err := engine.GetCreatorByAddress(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newCreator := models.Creator{WalletAddress: address.String()}
			creationErr := dbutils.DB.Create(&newCreator).Error
			if creationErr != nil {
				return nil, creationErr
			}

			err = CreateWalletAccount(newCreator.ID.String())
			if err != nil {
				return nil, err
			}
			return &newCreator, nil
		}
	}

	return cachedCreator, nil
}
