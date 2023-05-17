package internal

import "time"

// {
// 	"kid": "a46850a8-d024-4740-9f27-5de9c61a2182",
// 	"aud": "https://demo.dynamic.xyz",
// 	"iss": "app.dynamic.xyz/a46850a8-d024-4740-9f27-5de9c61a2182",
// 	"sub": "025887d9-bf4d-4f14-9e9a-5e919ed0908b",
// 	"sid": "266f77e0-2120-4aab-92eb-0f67f229179f",
// 	"alias": "",
// 	"environment_id": "a46850a8-d024-4740-9f27-5de9c61a2182",
// 	"lists": [],
// 	"missing_fields": [],
// 	"scope": "",
// 	"verified_credentials": [
// 	  {
// 		"address": "0x7500613737768c5EfD9411503D7a55E57287u817",
// 		"chain": "eip155",
// 		"id": "a764d9a2-b6ad-48a7-9837-15e9ff22ab98",
// 		"name_service": {},
// 		"public_identifier": "0x7500613737768c5EfD3F11503D7a55E57287d476",
// 		"wallet_name": "metamask",
// 		"wallet_provider": "browserExtension",
// 		"format": "blockchain"
// 	  },
// 	  {
// 		"id": "6640da56-a851-4bcb-bfd4-4bc39928036e",
// 		"public_identifier": "john.doe",
// 		"format": "oauth",
// 		"oauth_provider": "discord",
// 		"oauth_username": "john.doe",
// 		"oauth_display_name": null,
// 		"oauth_account_id": "961311833355092068",
// 		"oauth_account_photos": [],
// 		"oauth_emails": []
// 	  },
// 	  {
// 		"id": "6640da56-a851-4bcb-bfd4-4bc39928036e",
// 		"public_identifier": "test.user@gmail.com",
// 		"email": "test.user@gmail.com",
// 		"format": "email"
// 	  }
// 	],
// 	"last_verified_credential_id": "a764d9a2-b6ad-48a7-9837-15e9ff22ab98",
// 	"first_visit": "2023-03-05T00:11:30.459Z",
// 	"last_visit": "2023-04-21T17:18:33.114Z",
// 	"new_user": false,
// 	"iat": 1682097622,
// 	"exp": 1682104822
//   }

type DynamicJWTMetadata struct {
	Kid                      string                `json:"kid,omitempty"`
	Aud                      string                `json:"aud,omitempty"`
	Iss                      string                `json:"iss,omitempty"`
	Sub                      string                `json:"sub,omitempty"`
	Sid                      string                `json:"sid,omitempty"`
	Alias                    string                `json:"alias,omitempty"`
	EnvironmentID            string                `json:"environment_id,omitempty"`
	Lists                    []interface{}         `json:"lists,omitempty"`
	MissingFields            []interface{}         `json:"missing_fields,omitempty"`
	Scope                    string                `json:"scope,omitempty"`
	VerifiedCredentials      []VerifiedCredentials `json:"verified_credentials,omitempty"`
	LastVerifiedCredentialID string                `json:"last_verified_credential_id,omitempty"`
	FirstVisit               time.Time             `json:"first_visit,omitempty"`
	LastVisit                time.Time             `json:"last_visit,omitempty"`
	NewUser                  bool                  `json:"new_user,omitempty"`
	Iat                      int                   `json:"iat,omitempty"`
	Exp                      int                   `json:"exp,omitempty"`
}

type NameService struct {
}
type VerifiedCredentials struct {
	Address            string        `json:"address,omitempty"`
	Chain              string        `json:"chain,omitempty"`
	ID                 string        `json:"id,omitempty"`
	NameService        NameService   `json:"name_service,omitempty"`
	PublicIdentifier   string        `json:"public_identifier,omitempty"`
	WalletName         string        `json:"wallet_name,omitempty"`
	WalletProvider     string        `json:"wallet_provider,omitempty"`
	Format             string        `json:"format,omitempty"`
	OauthProvider      string        `json:"oauth_provider,omitempty"`
	OauthUsername      string        `json:"oauth_username,omitempty"`
	OauthDisplayName   interface{}   `json:"oauth_display_name,omitempty"`
	OauthAccountID     string        `json:"oauth_account_id,omitempty"`
	OauthAccountPhotos []interface{} `json:"oauth_account_photos,omitempty"`
	OauthEmails        []interface{} `json:"oauth_emails,omitempty"`
	Email              string        `json:"email,omitempty"`
}
