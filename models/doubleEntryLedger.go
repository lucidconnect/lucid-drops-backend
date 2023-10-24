package models

import (
	uuid "github.com/satori/go.uuid"
)

type DoubleEntryLedger struct {
	Base
	TransactionReference string `gorm:"not null" json:"transaction_reference"`
	SourceAccoountID     uuid.UUID
	DestinationAccountID uuid.UUID
	Amount               int64  `gorm:"not null" json:"amount"`
	TransactionType      string `gorm:"not null" json:"transaction_type"`
	//ID of the corresponding double entry row
	PartnerID *string `json:"partner_id"`
	LedgerID  *string `json:"ledger_id"`
}
