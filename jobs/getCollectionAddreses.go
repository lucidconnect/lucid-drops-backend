package jobs

import (
	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/models"
)

// func FillOutContractAddresses() {
// 	deployments, err := fetchAllMissingContractAddressesDeployments()
// 	if err != nil {
// 		log.Error().Msg(err.Error())
// 		return
// 	}

// 	for _, deployment := range deployments {
// 		contractAdddress, err := drops.GetOnchainContractAddressFromDeploymentHash(*deployment.AAWalletDeploymentHash)
// 		if err != nil {
// 			log.Error().Msg(err.Error())
// 			notifier.NotifyTelegram(fmt.Sprintf("👺 Drop deployment failed (%s)", err), structure.EngineeringTeam)
// 			continue
// 		}

// 		notifier.NotifyTelegram("🪼 New Drop deployed at "+contractAdddress, structure.EngineeringTeam)

// 		deployment.AAContractAddress = &contractAdddress

// 		err = engine.SaveModel(deployment)
// 		if err != nil {
// 			notifier.NotifyTelegram(fmt.Sprintf("👺 Drop (%s )Saving failed (%s)", contractAdddress, err), structure.EngineeringTeam)
// 			continue
// 		}
// 	}
// }

func fetchAllMissingContractAddressesDeployments() ([]models.Drop, error) {
	var drops []models.Drop
	err := dbutils.DB.Where("aa_wallet_deployment_hash IS NOT NULL AND aa_wallet_deployment_hash !='' AND aa_contract_address IS NULL").Find(&drops).Error
	if err != nil {
		return nil, err
	}

	return drops, nil
}
