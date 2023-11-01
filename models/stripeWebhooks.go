package models

type StripeWebhooks struct {
	Base
	RequestID      string
	IdempotencyKey string
	Processed      bool `gorm:"default:false"`
	ErrorMetaData  string
}
