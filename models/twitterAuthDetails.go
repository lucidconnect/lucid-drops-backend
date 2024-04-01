package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type TwitterAuthDetails struct {
	ID               uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt        time.Time      `gorm:"not null"`
	UpdatedAt        time.Time      `gorm:"not null"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	OAuthToken       string         `json:"oauth_token"`
	OAuthTokenSecret string         `json:"oauth_token_secret"`
	UserID           string         `gorm:"primaryKey" json:"user_id"`
	ScreenName       string         `json:"screen_name"`
	ItemID           *string        `gorm:"default:null;"`
	WhiteListed      bool           `gorm:"default:false;" json:"white_listed"`
}
