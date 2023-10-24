package models

import "inverse.so/graph/model"

type Wallet struct {
	Base
	CreatorID     string       `gorm:"type:uuid;index:idx_wallet_creatorId,unique;not null" json:"creator_id"`
	BalanceBase   int64        `gorm:"type:bigint;default:0" json:"balance_base"`
	CanBeNegative bool         `gorm:"default:false"`
	Currency      CurrencyType `gorm:"default:USD" json:"currency"`
}

type CurrencyType string

const (
	NGN CurrencyType = "NGN"
	USD CurrencyType = "USD"
	GBP CurrencyType = "GBP"
	EUR CurrencyType = "EUR"
)

func (c CurrencyType) String() string {
	return string(c)
}

func (w *Wallet) ToGraphData() *model.Wallet {
	return &model.Wallet{
		Balance:  int(w.BalanceBase),
		Currency: w.Currency.String(),
	}
}
