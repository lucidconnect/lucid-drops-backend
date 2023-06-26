package models

type TelegramAuthDetails struct {
	Base
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Hash      string `json:"hash"`
	PhotoURL  string `json:"photo_url"`
}
