package jobs

import (
	"fmt"

	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/engine/collections"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/notifier"
	"github.com/lucidconnect/inverse/structure"
	"github.com/rs/zerolog/log"
)

func FillOutContractAddresses() {
	deployments, err := fetchAllMissingContractAddressesDeployments()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	for _, deployment := range deployments {
		contractAdddress, err := collections.GetOnchainContractAddressFromDeploymentHash(*deployment.AAWalletDeploymentHash)
		if err != nil {
			log.Error().Msg(err.Error())
			notifier.NotifyTelegram(fmt.Sprintf("👺 Collection deployment failed (%s)", err), structure.EngineeringTeam)
			continue
		}

		notifier.NotifyTelegram("🪼 New Collection deployed at "+*contractAdddress, structure.EngineeringTeam)

		deployment.AAContractAddress = contractAdddress

		err = engine.SaveModel(deployment)
		if err != nil {
			notifier.NotifyTelegram(fmt.Sprintf("👺 Collection (%s )Saving failed (%s)", *contractAdddress, err), structure.EngineeringTeam)
			continue
		}
	}
}

func fetchAllMissingContractAddressesDeployments() ([]models.Collection, error) {
	var collections []models.Collection
	err := dbutils.DB.Where("aa_wallet_deployment_hash IS NOT NULL AND aa_wallet_deployment_hash !='' AND aa_contract_address IS NULL").Find(&collections).Error
	if err != nil {
		return nil, err
	}

	return collections, nil
}
