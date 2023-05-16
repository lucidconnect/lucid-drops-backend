package internal

import "time"

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
