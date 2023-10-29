package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type PrivyUserResponse struct {
	ID             string `json:"id"`
	CreatedAt      int    `json:"created_at"`
	LinkedAccounts []struct {
		Type             string `json:"type"`
		Address          string `json:"address"`
		VerifiedAt       int    `json:"verified_at"`
		ChainID          string `json:"chain_id,omitempty"`
		ChainType        string `json:"chain_type,omitempty"`
		WalletClient     string `json:"wallet_client,omitempty"`
		WalletClientType string `json:"wallet_client_type,omitempty"`
		ConnectorType    string `json:"connector_type,omitempty"`
		RecoveryMethod   string `json:"recovery_method,omitempty"`
	} `json:"linked_accounts"`
}

// FUTURE add caching
func GetPrivyAppIdAndSecret() (string, string) {
	return os.Getenv("PRIVY_APP_ID"), os.Getenv("PRIVY_APP_SECRET")
}

// https://docs.privy.io/guide/backend/api/users/get
func FetchPrivyUser(userDid string) (*PrivyUserResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://auth.privy.io/api/v1/users/%s", userDid), nil)
	if err != nil {
		return nil, err
	}

	privyAppId, privyAppSecret := GetPrivyAppIdAndSecret()

	req.SetBasicAuth(privyAppId, privyAppSecret)
	req.Header.Add("privy-app-id", privyAppId)
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Accept", "*/*")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response PrivyUserResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// format "did:privy:<userid>"
func GetPrivyWalletsFromSubKey(privySubKey string) (*common.Address, error) {
	parts := strings.Split(privySubKey, ":")
	if len(parts) != 3 {
		return nil, errors.New("invalid Privy Token Supplied")
	}

	userDid := parts[2]

	userResponse, err := FetchPrivyUser(userDid)
	if err != nil {
		return nil, errors.New("couldn't get privy user details")
	}

	var embededWallet *common.Address

	for _, accounts := range userResponse.LinkedAccounts {
		if accounts.Type == "wallet" && accounts.WalletClientType == "embedded" {
			parsedAddress := common.HexToAddress(accounts.Address)
			embededWallet = &parsedAddress
		}
	}

	if embededWallet == nil {
		return nil, errors.New("user lacks an embedded wallet")
	}

	return embededWallet, nil
}
