package engine

import (
	"errors"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
	"inverse.so/dbutils"
	"inverse.so/models"
	"inverse.so/utils"
)

func AttachContractAddressForCreationHash(transactionHash, contractAddress string) error {
	collection, err := GetCollectionByDeploymentHash(transactionHash)
	if err != nil {
		return err
	}

	collection.ContractAddress = utils.GetStrPtr(contractAddress)

	return SaveModel(collection)
}

func GetCollectionByDeploymentHash(deploymentHash string) (*models.Collection, error) {
	var collection models.Collection

	err := dbutils.DB.Where("transaction_hash=?", deploymentHash).First(&collection).Error
	if err != nil {
		return nil, errors.New("collection not found")
	}

	return &collection, nil
}

func GetCreatorByID(creatorID string) (*models.Creator, error) {
	var creator models.Creator

	err := dbutils.DB.Where("id=?", creatorID).First(&creator).Error
	if err != nil {
		return nil, errors.New("crator not found")
	}

	return &creator, nil
}

func GetCreatorByAddress(address string) (*models.Creator, error) {
	var creator models.Creator

	err := dbutils.DB.Where(&models.Creator{WalletAddress: address}).First(&creator).Error
	if err != nil {
		return nil, errors.New("address not found")
	}

	return &creator, nil
}

func GetCreatorByInverseUsername(inverseUsername string) (*models.Creator, error) {
	var creator models.Creator

	err := dbutils.DB.Where(&models.Creator{InverseUsername: &inverseUsername}).First(&creator).Error
	if err != nil {
		return nil, errors.New("username isn't being used")
	}

	return &creator, nil
}

func GetCollectionByID(collectionID string) (*models.Collection, error) {
	var collection models.Collection

	err := dbutils.DB.Where("id=?", collectionID).First(&collection).Error
	if err != nil {
		return nil, errors.New("collection not found")
	}

	return &collection, nil
}

func GetItemByID(itemID string) (*models.Item, error) {
	var item models.Item

	err := dbutils.DB.Preload(clause.Associations).Where("id=?", itemID).First(&item).Error
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

	err := dbutils.DB.Where(&models.EmailDomainWhiteList{
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

	err := dbutils.DB.Where(&models.SingleEmailClaim{
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

	err := dbutils.DB.Where("creator_id=?", creatorID).Find(&collections).Error
	if err != nil {
		return nil, errors.New("collections not found")
	}

	return collections, nil
}

func GetCollectionItems(collectionID string) ([]*models.Item, error) {
	var items []*models.Item

	err := dbutils.DB.Where("collection_id=?", collectionID).Find(&items).Error
	if err != nil {
		return nil, errors.New("items not found")
	}

	return items, nil
}

func GetAuthorizedSubdomainsForItem(itemID string) ([]*models.EmailDomainWhiteList, error) {
	var subDomains []*models.EmailDomainWhiteList

	err := dbutils.DB.Where("item_id=?", itemID).Find(&subDomains).Error
	if err != nil {
		return nil, errors.New("subdomains not found")
	}

	return subDomains, nil
}

func GetEmailOTPRecordByID(recordID string) (*models.EmailOTP, error) {
	var emailOTP models.EmailOTP

	err := dbutils.DB.Where("id=?", recordID).First(&emailOTP).Error
	if err != nil {
		return nil, errors.New("email verification not found")
	}

	return &emailOTP, nil
}

func FetchTwitterAuthByID(authID string) (*models.TwitterAuthDetails, error) {

	var twitterAuth models.TwitterAuthDetails
	err := dbutils.DB.Where("id=?", authID).First(&twitterAuth).Error
	if err != nil {
		return nil, errors.New("twitter auth not found")
	}

	return &twitterAuth, nil
}

func FetchTelegramAuthByID(authID string) (*models.TelegramAuthDetails, error) {

	var telegramAuth models.TelegramAuthDetails
	err := dbutils.DB.Where("id=?", authID).First(&telegramAuth).Error
	if err != nil {
		return nil, errors.New("twitter auth not found")
	}

	return &telegramAuth, nil
}

func FetchTelegramCriteriaByLink(channelLink string) (*models.TelegramCriteria, error) {
	var criteria models.TelegramCriteria

	err := dbutils.DB.Where("channel_link=?", channelLink).First(&criteria).Error
	if err != nil {
		return nil, errors.New("telegram criteria not found")
	}

	return &criteria, nil
}

func CreateModel(newModel interface{}) error {
	return dbutils.DB.Create(newModel).Error
}

func SaveModel(model interface{}) error {
	return dbutils.DB.Save(model).Error
}
