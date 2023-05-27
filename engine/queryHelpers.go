package engine

import (
	"errors"

	"inverse.so/models"
	"inverse.so/utils"
)

func GetCreatorByAddress(address string) (*models.Creator, error) {
	var creator models.Creator

	err := utils.DB.Where(&models.Creator{WalletAddress: address}).First(&creator).Error
	if err != nil {
		return nil, errors.New("address not found")
	}

	return &creator, nil
}

func GetCreatorByInverseUsername(inverseUsername string) (*models.Creator, error) {
	var creator models.Creator

	err := utils.DB.Where(&models.Creator{InverseUsername: &inverseUsername}).First(&creator).Error
	if err != nil {
		return nil, errors.New("username isn't being used")
	}

	return &creator, nil
}

func GetCollectionByID(collectionID string) (*models.Collection, error) {
	var collection models.Collection

	err := utils.DB.Where("id=?", collectionID).First(&collection).Error
	if err != nil {
		return nil, errors.New("collection not found")
	}

	return &collection, nil
}

func GetCreatorCollections(creatorID string) ([]*models.Collection, error) {
	var collections []*models.Collection

	err := utils.DB.Where("creator_id=?", creatorID).Find(&collections).Error
	if err != nil {
		return nil, errors.New("collection not found")
	}

	return collections, nil
}

func CreateCollection(newCollection *models.Collection) error {
	return utils.DB.Create(newCollection).Error
}

func SaveCollection(collection *models.Collection) error {
	return utils.DB.Save(collection).Error
}
