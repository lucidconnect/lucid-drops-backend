package models

type PatreonAuthDetails struct {
	Base
	Code        string `json:"code"`
	AccessToken string `json:"access_token"`
}
