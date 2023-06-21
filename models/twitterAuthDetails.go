package models

type TwitterAuthDetails struct {
	Base
	OAuthToken       string `json:"oauth_token"`
	OAuthTokenSecret string `json:"oauth_token_secret"`
	UserID           string `gorm:"primaryKey" json:"user_id"`
	ScreenName       string `json:"screen_name"`
}