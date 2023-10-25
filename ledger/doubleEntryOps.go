package ledger

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"

	"inverse.so/models"
	"inverse.so/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrAccountNotActive = errors.New("account not active")
var ErrInsufficientBalance = errors.New("insufficient balance")
var ErrAccountCannotBeNegative = errors.New("account cannot be negative")

func (l *Ledger) Transfer(tx *gorm.DB, instruction TransferInstruction) error {

	isLocalTx := tx == nil
	if isLocalTx {
		tx = l.DB.Begin()
	}

	//fetch and lock accounts
	accounts, err := l.fetchAndLockTransferScope(tx, instruction)
	if err != nil {
		return err
	}

	if isNegativeInt(instruction.Amount) {
		tx.Rollback()
		return errors.New("ledger instruction amount cannot be negative")
	}

	var sourceAccount models.Wallet
	var destinationAccount models.Wallet
	var creditID = fmt.Sprintf("%s-%s", instruction.TxRef, utils.RandUpperCaseAlphaNumericRunes(5))
	var debitID = fmt.Sprintf("%s-%s", instruction.TxRef, utils.RandUpperCaseAlphaNumericRunes(5))

	switch instruction.Side {
	case Debit:
		destinationAccount = accounts[0]
		sourceAccount = accounts[1]
		if !accounts[0].CanBeNegative {
			destinationAccount = accounts[1]
			sourceAccount = accounts[0]
		}

		if !sourceAccount.CanBeNegative && isGreaterThanBalance(instruction.Amount, sourceAccount.BalanceBase) {
			tx.Rollback()
			return ErrInsufficientBalance
		}
	case Credit:
		destinationAccount = accounts[1]
		sourceAccount = accounts[0]
		if !accounts[0].CanBeNegative {
			destinationAccount = accounts[0]
			sourceAccount = accounts[1]
		}
	}

	creditSideLedgeEntry := models.DoubleEntryLedger{
		TransactionReference: instruction.TxRef,
		SourceAccoountID:     sourceAccount.ID,
		DestinationAccountID: destinationAccount.ID,
		Amount:               instruction.Amount,
		TransactionType:      Credit.String(),
		LedgerID:             &creditID,
		PartnerID:            &debitID,
	}

	debitSideLedgeEntry := models.DoubleEntryLedger{
		TransactionReference: instruction.TxRef,
		SourceAccoountID:     destinationAccount.ID,
		DestinationAccountID: sourceAccount.ID,
		Amount:               toNegativeInt(instruction.Amount),
		TransactionType:      Debit.String(),
		LedgerID:             &debitID,
		PartnerID:            &creditID,
	}

	//create ledger entries

	entries := []models.DoubleEntryLedger{creditSideLedgeEntry, debitSideLedgeEntry}
	for _, entry := range entries {
		err := tx.Create(&entry).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if !sourceAccount.CanBeNegative {
		if isNegativeInt(sourceAccount.BalanceBase) {
			tx.Rollback()
			return ErrInsufficientBalance
		}
	}

	//update account balances
	err = tx.Model(&sourceAccount).Update("balance_base", gorm.Expr("balance_base + ?", toNegativeInt(instruction.Amount))).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&destinationAccount).Update("balance_base", gorm.Expr("balance_base + ?", instruction.Amount)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if isLocalTx {
		return tx.Commit().Error
	}

	return nil
}

func (l *Ledger) fetchAccountsByCreatorID(creatorID string) (*models.Wallet, error) {

	var accounts models.Wallet
	err := l.DB.Where("creator_id = ?", creatorID).First(&accounts).Error
	if err != nil {
		return nil, err
	}

	return &accounts, nil
}

func (l *Ledger) fetchSysAccount() (*models.Wallet, error) {
	var sysAccount models.Wallet
	err := l.DB.Model(&models.Wallet{}).Where("can_be_negative = ?", true).First(&sysAccount).Error
	if err != nil {
		return nil, err
	}

	log.Print(sysAccount)
	return &sysAccount, nil
}

func (l *Ledger) fetchAndLockTransferScope(tx *gorm.DB, instruction TransferInstruction) ([]models.Wallet, error) {

	accounts := make([]models.Wallet, 2)
	sysAccount, err := l.fetchSysAccount()
	if err != nil {
		return nil, err
	}

	accounts[0] = *sysAccount

	userAccount, err := l.fetchAccountsByCreatorID(instruction.UserID.String())
	if err != nil {
		return nil, err
	}

	accounts[1] = *userAccount

	err = l.lockTransferScope(tx, &accounts)
	if err != nil {
		return nil, err
	}

	log.Print(accounts)
	return accounts, nil
}

func (l *Ledger) lockTransferScope(tx *gorm.DB, accounts *[]models.Wallet) error {

	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id IN (?)", []uuid.UUID{(*accounts)[0].ID, (*accounts)[1].ID}).
		Find(&accounts).Error
	if err != nil {
		return err
	}

	return nil
}

func isNegativeInt(value int64) bool {

	return value < 0

}

func toNegativeInt(value int64) int64 {

	if value < 0 {
		return value
	}
	return value * -1
}

func isGreaterThanBalance(value int64, balance int64) bool {
	return value > balance
}
