package whitelist

import (
	"errors"
	"time"

	"gorm.io/gorm/clause"
	"inverse.so/dbutils"
	"inverse.so/emails"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
	"inverse.so/utils"
)

func sendEmailOnCreate(dbEmail *models.SingleEmailClaim) error {
	from := "noreply@getabacus.app"

	item, err := engine.GetItemByID(dbEmail.ItemID.String())
	if err != nil {
		return err
	}

	creator, err := engine.GetCreatorByID(dbEmail.CreatorID.String())
	if err != nil {
		return err
	}

	mintPass, err := CreateMintPassForNoCriteriaItem(dbEmail.ItemID.String())
	if err != nil {
		return err
	}

	claimLink := utils.UseEnvOrDefault("FE_BASE_URL", "https://inverse.so") + "/claim/" + item.ID.String() + "?requestId=" + *mintPass.PassID
	err = emails.SendClaimNudgeEmail(dbEmail.EmailAddress, from, item.Name, claimLink, *creator.InverseUsername, item.Image)
	if err != nil {
		return err
	}
	return nil
}

func CreateEmailWhitelistForItem(input *model.NewEmailWhitelistInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new collection")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if item.Criteria != nil {
		//Delete Existing criteria
		err := engine.DeleteCriteriaIfExists(item)
		if err != nil {
			return nil, err
		}
	}

	if len(input.AuthorizedEmails) == 0 {
		return nil, errors.New("please passing in email list")
	}

	dbEmails := make([]*models.SingleEmailClaim, len(input.AuthorizedEmails))
	for idx, email := range input.AuthorizedEmails {
		dbEmails[idx] = &models.SingleEmailClaim{
			CreatorID:    creator.ID,
			ItemID:       item.ID,
			EmailAddress: email,
		}
	}

	insertionErr := dbutils.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dbEmails, 100).Error
	if insertionErr != nil {
		return nil, insertionErr
	}

	for _, dbEmail := range dbEmails {
		if err = sendEmailOnCreate(dbEmail); err != nil {
			continue
		}

		timeNow := time.Now()

		dbEmail.SentOutAt = &timeNow

		insertionErr := dbutils.DB.Save(dbEmail).Error

		if insertionErr != nil {
			continue
		}
	}

	emailCriteria := model.ClaimCriteriaTypeEmailWhiteList
	item.Criteria = &emailCriteria
	itemUpdateErr := engine.SaveModel(nil, item)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	return item.ToGraphData(), nil
}

func FetchCriteriaAuthorizedEmails(itemID string) ([]string, error) {
	var dbEmails []*models.SingleEmailClaim
	err := dbutils.DB.Where("item_id = ?", itemID).Find(&dbEmails).Error
	if err != nil {
		return nil, err
	}

	emails := make([]string, len(dbEmails))
	for idx, dbEmail := range dbEmails {
		emails[idx] = dbEmail.EmailAddress
	}

	return emails, nil
}
