package models

type TwitterAuthDetails struct {
	Base
	OAuthToken       string  `json:"oauth_token"`
	OAuthTokenSecret string  `json:"oauth_token_secret"`
	UserID           string  `gorm:"primaryKey" json:"user_id"`
	ScreenName       string  `json:"screen_name"`
	ItemID           *string `gorm:"default:null;"`
	WhiteListed      bool    `gorm:"default:false;" json:"white_listed"`
}
