package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type WalletAddressClaim struct {
	BaseWithoutPrimaryKey
	CreatorID             uuid.UUID `gorm:"primaryKey"`
	DropID                uuid.UUID `gorm:"primaryKey"`
	WalletAddress         string    `gorm:"primaryKey"`
	ENS                   *string
	EmbeddedWalletAddress string     `gorm:"default:null"`
	SentOutAt             *time.Time `gorm:"default:null"`
}
