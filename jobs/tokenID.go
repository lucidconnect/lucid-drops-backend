package jobs

import (
	"inverse.so/dbutils"
	"inverse.so/models"
)



func fetchItemsWithUnresolvedTokenIDs() *[]models.Item {
	var items []models.Item
	err := dbutils.DB.Where("token_id IS NULL").Find(&items).Error
	if err != nil {
		return nil
	}

	return &items
}
