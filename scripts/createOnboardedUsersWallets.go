package scripts

import (
	"log"

	"inverse.so/dbutils"
	"inverse.so/models"
)

func CreateOnboardedUsersWallets() {

	var creators []models.Creator
	dbutils.DB.Where("wallet_address IS NOT NULL").Find(&creators)
	for _, creator := range creators {
		err := dbutils.DB.Model(&models.Wallet{}).Where("creator_id = ?", creator.ID).FirstOrCreate(&models.Wallet{
			CreatorID: creator.ID.String(),
		}).Error
		if err != nil {
			log.Printf("Error creating wallet for creator %v\n", err)
		}
	}
	
}
