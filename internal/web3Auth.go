package internal

type Web3AuthMetadata struct {
	Iat     int    `json:"iat"`
	Aud     string `json:"aud"`
	Nonce   string `json:"nonce"`
	Iss     string `json:"iss"`
	Wallets []struct {
		PublicKey string `json:"public_key"`
		Type      string `json:"type"`
		Curve     string `json:"curve"`
		Address string `json:"address"`
	} `json:"wallets"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Verifier          string `json:"verifier"`
	VerifierID        string `json:"verifierId"`
	AggregateVerifier string `json:"aggregateVerifier"`
	Exp               int    `json:"exp"`
}
