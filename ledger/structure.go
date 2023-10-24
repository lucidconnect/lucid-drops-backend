package ledger

import uuid "github.com/satori/go.uuid"

type TransferInstruction struct {
	UserID uuid.UUID
	Amount int64
	TxRef  string
	Side   TransactionSide
}

type TransactionSide string

func (t TransactionSide) String() string {
	return string(t)
}

const (
	Debit  TransactionSide = "debit"
	Credit TransactionSide = "credit"
)
