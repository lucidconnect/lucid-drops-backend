package whitelist

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/ledger"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func IsThisAValidEthAddress(maybeAddress string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	if len(maybeAddress) != 43 {
		return false
	}

	return re.MatchString(maybeAddress)
}

func CreateMintPassForNoCriteriaItem(itemID string) (*model.ValidationRespoonse, error) {

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, err
	}

	if item.ClaimDeadline != nil {
		if time.Now().After(*item.ClaimDeadline) {
			return nil, errors.New("the item is no longer available to be claimed")
		}
	}

	// if item.Criteria == nil || *item.Criteria != model.ClaimCriteriaTypeEmptyCriteria {
	// 	return nil, errors.New("unable to generate mintpass for this item")
	// }

	drop, err := engine.GetDropByID(item.DropID.String())
	if err != nil {
		return nil, errors.New("drop not found")
	}

	if drop.AAContractAddress == nil {
		return nil, errors.New("drop contract address not found")
	}

	if item.TokenID == nil {
		return nil, errors.New("The requested item is not ready to be claimed, please try again in a few minutes")
	}

	if ItemOverEditionLimit(item) {
		return nil, errors.New("item edition limit reached")
	}

	tx := dbutils.DB.Begin()
	newMint := models.MintPass{
		ItemId:              item.ID.String(),
		ItemIdOnContract:    *item.TokenID,
		DropContractAddress: *drop.AAContractAddress,
		BlockchainNetwork:   drop.BlockchainNetwork,
	}

	err = tx.Create(&newMint).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return &model.ValidationRespoonse{
		Valid:  true,
		PassID: utils.GetStrPtr(newMint.ID.String()),
	}, nil
}

func chargeClaimFee(userID string, item *models.Item, tx *gorm.DB) error {

	inverseMargin := 0.25
	marginDeduction := int64(float64(item.ClaimFee) * inverseMargin)
	claimFeeAfterMarginDeduction := int64(item.ClaimFee) - marginDeduction

	drop, err := engine.GetDropByID(item.DropID.String())
	if err != nil {
		return errors.New("drop not found")
	}

	///debit side instruction for collector
	l := ledger.New(dbutils.DB)
	debitInstruction := ledger.TransferInstruction{
		UserID: uuid.FromStringOrNil(userID),
		Amount: int64(item.ClaimFee),
		Side:   ledger.Debit,
		TxRef:  fmt.Sprintf("claim-%s-%s-%s", item.ID.String(), userID, utils.RandAlphaNumericRunes(5)),
	}

	err = l.Transfer(tx, debitInstruction)
	if err != nil {
		return err
	}

	//credit side instruction for creator
	creditInstruction := ledger.TransferInstruction{
		UserID: uuid.FromStringOrNil(drop.CreatorID.String()),
		Amount: claimFeeAfterMarginDeduction,
		Side:   ledger.Credit,
		TxRef:  fmt.Sprintf("claim-%s-%s-%s", item.ID.String(), drop.CreatorID.String(), utils.RandAlphaNumericRunes(5)),
	}

	err = l.Transfer(tx, creditInstruction)
	if err != nil {
		return err
	}

	// send creator successful charge/money made email

	//credit side instruction for inverse margin
	collectInstruction := ledger.TransferInstruction{
		UserID: uuid.FromStringOrNil(l.CollectAccount.CreatorID),
		Amount: marginDeduction,
		Side:   ledger.Credit,
		TxRef:  fmt.Sprintf("claim-%s-%s-%s", item.ID.String(), drop.CreatorID.String(), utils.RandAlphaNumericRunes(5)),
	}

	err = l.Transfer(tx, collectInstruction)
	if err != nil {
		return err
	}

	return nil
}

func CreateMintPassForValidatedCriteriaItem(itemID string) (*model.ValidationRespoonse, error) {

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, err
	}

	if item.ClaimDeadline != nil {
		if time.Now().After(*item.ClaimDeadline) {
			return nil, errors.New("the item is no longer available to be claimed")
		}
	}

	if item.Criteria == nil {
		return nil, errors.New("unable to generate mintpass for this item")
	}

	drop, err := engine.GetDropByID(item.DropID.String())
	if err != nil {
		return nil, errors.New("drop not found")
	}

	if drop.AAContractAddress == nil {
		return nil, errors.New("drop contract address not found")
	}

	if item.TokenID == nil {
		return nil, errors.New("The requested item is not ready to be claimed, please try again in a few minutes")
	}

	if ItemOverEditionLimit(item) {
		return nil, errors.New("item edition limit reached")
	}

	newMint := models.MintPass{
		ItemId:              item.ID.String(),
		ItemIdOnContract:    *item.TokenID,
		DropContractAddress: *drop.AAContractAddress,
		BlockchainNetwork:   drop.BlockchainNetwork,
	}

	err = dbutils.DB.Create(&newMint).Error
	if err != nil {
		return nil, err
	}

	return &model.ValidationRespoonse{
		Valid:  true,
		PassID: utils.GetStrPtr(newMint.ID.String()),
	}, nil
}

func ItemOverEditionLimit(item *models.Item) bool {

	if item.EditionLimit != nil {
		var editionCount int64
		err := dbutils.DB.Model(&models.MintPass{}).Where("item_id = ?", item.ID).Count(&editionCount).Error
		if err == nil {
			return int(editionCount) >= *item.EditionLimit
		}
	}

	return false
}
