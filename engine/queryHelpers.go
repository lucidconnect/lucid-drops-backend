package engine

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
	"inverse.so/dbutils"
	"inverse.so/graph/model"
	"inverse.so/models"
	"inverse.so/utils"
)

func AttachContractAddressForCreationHash(transactionHash, contractAddress string) error {
	collection, err := GetCollectionByDeploymentHash(transactionHash)
	if err != nil {
		return err
	}

	collection.AAContractAddress = utils.GetStrPtr(contractAddress)

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
		return nil, fmt.Errorf("creator (%s) not found", creatorID)
	}

	return &creator, nil
}

func GetMintPassById(passId string) (*models.MintPass, error) {
	var pass models.MintPass

	err := dbutils.DB.Where("id=?", passId).First(&pass).Error
	if err != nil {
		return nil, errors.New("mint pass not found")
	}

	return &pass, nil
}

func GetCreatorByAddress(address common.Address) (*models.Creator, error) {
	var creator models.Creator

	query := fmt.Sprintf("SELECT * FROM creators WHERE LOWER(wallet_address)='%s'", strings.ToLower(address.String()))
	err := dbutils.DB.Raw(query).First(&creator).Error
	if err != nil {
		return nil, err
	}

	return &creator, nil
}

func GetAltSignerByAddress(address common.Address) (*models.SignerInfo, error) {
	var altSigner models.SignerInfo

	err := dbutils.DB.Where("wallet_address=?", address.String()).First(&altSigner).Error
	if err != nil {
		return nil, fmt.Errorf("address (%s) not found", address)
	}

	return &altSigner, nil
}

func GetAltSignerByCreatorID(id string) (*models.SignerInfo, error) {
	var altSigner models.SignerInfo

	err := dbutils.DB.Where("creator_id=?", id).First(&altSigner).Error
	if err != nil {
		return nil, fmt.Errorf("creator signer (%s) not found", id)
	}

	return &altSigner, nil
}

func GetCreatorByInverseUsername(inverseUsername string) (*models.Creator, error) {
	var creator models.Creator

	err := dbutils.DB.Where("inverse_username=?", inverseUsername).First(&creator).Error
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

func GetClaimedItemByAddress(address string) ([]*models.Item, error) {
	var mintPasses []models.MintPass

	err := dbutils.DB.Where("minter_address=?", address).Find(&mintPasses).Error
	if err != nil {
		return nil, errors.New("collection not found")
	}

	itemsIds := make([]string, len(mintPasses))

	for idx, pass := range mintPasses {
		itemsIds[idx] = pass.ItemId
	}

	var claimedItems []*models.Item
	err = dbutils.DB.Where("id IN (?)", itemsIds).Find(&claimedItems).Error
	if err != nil {
		return nil, errors.New("claimed items no longer exists")
	}

	return claimedItems, nil
}

func GetItemByID(itemID string) (*models.Item, error) {
	var item models.Item

	err := dbutils.DB.Preload(clause.Associations).Where("id=?", itemID).First(&item).Error
	if err != nil {
		return nil, errors.New("item not found")
	}

	return &item, nil
}

func GetItemQuestionsByItem(item *models.Item) ([]*model.QuestionnaireType, error) {
	if item.Criteria == nil {
		return []*model.QuestionnaireType{}, nil
	}

	switch *item.Criteria {
	case model.ClaimCriteriaTypeDirectAnswerQuestionnaire:
		var directQuestions []*models.DirectAnswerCriteria

		err := dbutils.DB.Where(&models.DirectAnswerCriteria{ItemID: item.ID}).Find(&directQuestions).Error
		if err != nil {
			return nil, errors.New("seems item doesn't have any direct questions")
		}

		marshalledQuestion := make([]*model.QuestionnaireType, len(directQuestions))
		for idx, q := range directQuestions {
			marshalledQuestion[idx] = q.ToGraphData()
		}

		return marshalledQuestion, nil

	case model.ClaimCriteriaTypeMutliChoiceQuestionnaire:
		var multiChoiceQuestions []*models.MultiChoiceCriteria

		err := dbutils.DB.Where(&models.MultiChoiceCriteria{ItemID: item.ID}).Find(&multiChoiceQuestions).Error
		if err != nil {
			return nil, errors.New("seems item doesn't have any multi choice questions")
		}

		marshalledQuestion := make([]*model.QuestionnaireType, len(multiChoiceQuestions))
		for idx, q := range multiChoiceQuestions {
			marshalledQuestion[idx] = q.ToGraphData()
		}

		return marshalledQuestion, nil

	default:
		return nil, fmt.Errorf("item has a claim criteria of (%s) which doesn't have emais", *item.Criteria)
	}

}

func GetEmailClaimIDByItemAndEmailSubDomain(itemID *uuid.UUID, emailAddress string) (*models.EmailDomainWhiteList, error) {
	emailParts := strings.Split(emailAddress, "@")
	if len(emailParts) != 2 {
		return nil, fmt.Errorf("(%s) is not a valid email", emailAddress)
	}

	var claim models.EmailDomainWhiteList

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
	err := dbutils.DB.Model(&models.TelegramAuthDetails{}).Where("id=?", authID).First(&telegramAuth).Error
	if err != nil {
		return nil, errors.New("twitter auth not found")
	}

	return &telegramAuth, nil
}

func FetchPatreonAuthByID(authID string) (*models.PatreonAuthDetails, error) {

	var patreonAuth models.PatreonAuthDetails
	err := dbutils.DB.Model(&models.PatreonAuthDetails{}).Where("id=?", authID).First(&patreonAuth).Error
	if err != nil {
		return nil, errors.New("patreon auth not found")
	}

	return &patreonAuth, nil
}

func FetchTelegramCriteriaByLink(channelLink string) (*models.TelegramCriteria, error) {
	var criteria models.TelegramCriteria

	err := dbutils.DB.Where("channel_link=?", channelLink).First(&criteria).Error
	if err != nil {
		return nil, errors.New("telegram criteria not found")
	}

	return &criteria, nil
}

func GetFeaturedItems() ([]*models.Item, error) {
	var items []*models.Item

	err := dbutils.DB.Where("featured=?", true).Find(&items).Error
	if err != nil {
		return nil, errors.New("items not found")
	}

	return items, nil
}

func GetFeaturedCollections() ([]*models.Collection, error) {
	var collections []*models.Collection

	err := dbutils.DB.Where("featured=?", true).Find(&collections).Error
	if err != nil {
		return nil, errors.New("collections not found")
	}

	return collections, nil
}

func DeleteCriteriaIfExists(item *models.Item) error {
	var err error
	if item.Criteria == nil {
		switch *item.Criteria {
		case model.ClaimCriteriaTypeDirectAnswerQuestionnaire:
			//Delete existing questionnaire criteria
			err = dbutils.DB.Delete(&models.DirectAnswerCriteria{}, "item_id = ?", item.ID).Error
			if err != nil {
				return errors.New("an error occured while updating updating questionnaire criteria")
			}

			err = dbutils.DB.Delete(&models.MultiChoiceCriteria{}, "item_id = ?", item.ID).Error
			if err != nil {
				return errors.New("an error occured while updating updating questionnaire criteria")
			}
		case model.ClaimCriteriaTypeTwitterInteractions:
			//Delete existing twitter criteria
			err = dbutils.DB.Delete(&models.TwitterCriteria{}, "item_id = ?", item.ID).Error
			if err != nil {
				return errors.New("an error occured while updating updating twitter criteria")
			}
		case model.ClaimCriteriaTypeTwitterFollowers:
			//Delete existing twitter criteria
			err = dbutils.DB.Delete(&models.TwitterCriteria{}, "item_id = ?", item.ID).Error
			if err != nil {
				return errors.New("an error occured while updating updating twitter criteria")
			}
		case model.ClaimCriteriaTypePatreon:
			//Delete existing patreon criteria
			err = dbutils.DB.Delete(&models.PatreonCriteria{}, "item_id = ?", item.ID).Error
			if err != nil {
				return errors.New("an error occured while updating updating patreon criteria")
			}
		case model.ClaimCriteriaTypeTelegram:
			//Delete existing telegram criteria
			err = dbutils.DB.Delete(&models.TelegramCriteria{}, "item_id = ?", item.ID).Error
			if err != nil {
				return errors.New("an error occured while updating updating telegram criteria")
			}
		case model.ClaimCriteriaTypeEmailDomain:
			//Delete existing email domain criteria
			err = dbutils.DB.Delete(&models.EmailDomainWhiteList{}, "item_id = ?", item.ID).Error
			if err != nil {
				return errors.New("an error occured while updating updating email domain criteria")
			}
		case model.ClaimCriteriaTypeEmailWhiteList:
			//Delete existing email domain criteria
			err = dbutils.DB.Delete(&models.SingleEmailClaim{}, "item_id = ?", item.ID).Error
			if err != nil {
				return errors.New("an error occured while updating updating email domain criteria")
			}
		}
	}

	return nil
}

func GetUserProfileDetails(userName string) (*model.UserProfileType, *string, error) {
	var profile models.Creator
	err := dbutils.DB.Model(&models.Creator{}).Where("inverse_username=?", userName).First(&profile).Error
	if err != nil {
		return nil, nil, errors.New("user profile not found")
	}

	profileGraphData := profile.CreatorToProfileData()
	return profileGraphData, &profile.WalletAddress, nil
}

func CreateModel(newModel interface{}) error {
	return dbutils.DB.Create(newModel).Error
}

func SaveModel(model interface{}) error {
	return dbutils.DB.Save(model).Error
}

func SoftDeleteModel(model interface{}) error {
	return dbutils.DB.Delete(model).Error
}
