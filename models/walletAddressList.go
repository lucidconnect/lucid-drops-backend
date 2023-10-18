package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type WalletAddressClaim struct {
	BaseWithoutPrimaryKey
	CreatorID     uuid.UUID  `gorm:"primaryKey"`
	ItemID        uuid.UUID  `gorm:"primaryKey"`
	WalletAddress string     `gorm:"primaryKey"`
	SentOutAt     *time.Time `gorm:"default:null"`
}
