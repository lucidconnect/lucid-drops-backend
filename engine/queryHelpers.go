package engine

import (
	"errors"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
	"inverse.so/models"
	"inverse.so/utils"
)

func GetCreatorByID(creatorID string) (*models.Creator, error) {
	var creator models.Creator

	err := utils.DB.Where("id=?", creatorID).First(&creator).Error
	if err != nil {
		return nil, errors.New("crator not found")
	}

	return &creator, nil
}

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

func GetItemByID(itemID string) (*models.Item, error) {
	var item models.Item

	err := utils.DB.Where("id=?", itemID).First(&item).Error
	if err != nil {
		return nil, errors.New("item not found")
	}

	return &item, nil
}

func GetEmailClaimIDByItemAndEmailSubDomain(itemID *uuid.UUID, emailAddress string) (*models.SingleEmailClaim, error) {
	emailParts := strings.Split(emailAddress, "@")
	if len(emailParts) != 2 {
		return nil, fmt.Errorf("(%s) is not a valid email", emailAddress)
	}

	var claim models.SingleEmailClaim

	err := utils.DB.Where(&models.EmailDomainWhiteList{
		ItemID:     *itemID,
		BaseDomain: emailParts[1],
	}).First(&claim).Error

	if err != nil {
		return nil, errors.New("claim not found")
	}

	return &claim, nil
}

func GetEmailClaimIDByItemAndEmail(itemID *uuid.UUID, claimingEmail string) (*models.SingleEmailClaim, error) {
	var claim models.SingleEmailClaim

	err := utils.DB.Where(&models.SingleEmailClaim{
		ItemID:       *itemID,
		EmailAddress: claimingEmail,
	}).First(&claim).Error

	if err != nil {
		return nil, errors.New("claim not found")
	}

	return &claim, nil
}

func GetCreatorCollections(creatorID string) ([]*models.Collection, error) {
	var collections []*models.Collection

	err := utils.DB.Where("creator_id=?", creatorID).Find(&collections).Error
	if err != nil {
		return nil, errors.New("collections not found")
	}

	return collections, nil
}

func GetCollectionItems(collectionID string) ([]*models.Item, error) {
	var items []*models.Item

	err := utils.DB.Where("collection_id=?", collectionID).Find(&items).Error
	if err != nil {
		return nil, errors.New("items not found")
	}

	return items, nil
}

func GetAuthorizedSubdomainsForItem(itemID string) ([]*models.EmailDomainWhiteList, error) {
	var subDomains []*models.EmailDomainWhiteList

	err := utils.DB.Where("item_id=?", itemID).Find(&subDomains).Error
	if err != nil {
		return nil, errors.New("subdomains not found")
	}

	return subDomains, nil
}

func GetEmailOTPRecordByID(recordID string) (*models.EmailOTP, error) {
	var emailOTP models.EmailOTP

	err := utils.DB.Where("id=?", recordID).First(&emailOTP).Error
	if err != nil {
		return nil, errors.New("email verification not found")
	}

	return &emailOTP, nil
}

func CreateModel(newModel interface{}) error {
	return utils.DB.Create(newModel).Error
}

func SaveModel(model interface{}) error {
	return utils.DB.Save(model).Error
}
