package models

import (
	"time"
)

type PatreonAuthDetails struct {
	Base
	Code           string    `json:"code"`
	AccessToken    string    `json:"access_token"`
	RefreshToken   string    `json:"refresh_token"`
	UserID         string    `json:"user_id"`
	MembershipUIDs string    `json:"uid"`
	CampaignID     string    `json:"campaign_id"`
	ExpiresAt      time.Time `json:"expires_at"`
	WhiteListed    bool      `gorm:"default:false;" json:"white_listed"`
}
