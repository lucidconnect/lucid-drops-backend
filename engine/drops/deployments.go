package drops

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lucidconnect/inverse/addresswatcher"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/rs/zerolog/log"
)

type JiffyscanResponse struct {
	UserOps []struct {
		VerificationGasLimit string   `json:"verificationGasLimit"`
		UserOpHash           string   `json:"userOpHash"`
		TransactionHash      string   `json:"transactionHash"`
		Target               []string `json:"target"`
		AccountTarget        struct {
			Factory any `json:"factory"`
		} `json:"accountTarget"`
		Success       bool   `json:"success"`
		Signature     string `json:"signature"`
		Sender        string `json:"sender"`
		AccountSender struct {
			Factory string `json:"factory"`
		} `json:"accountSender"`
		RevertReason         any      `json:"revertReason"`
		PreVerificationGas   string   `json:"preVerificationGas"`
		PaymasterAndData     string   `json:"paymasterAndData"`
		Paymaster            string   `json:"paymaster"`
		Nonce                string   `json:"nonce"`
		Network              string   `json:"network"`
		MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
		MaxFeePerGas         string   `json:"maxFeePerGas"`
		Input                string   `json:"input"`
		GasPrice             string   `json:"gasPrice"`
		ID                   string   `json:"id"`
		GasLimit             string   `json:"gasLimit"`
		Factory              string   `json:"factory"`
		CallGasLimit         string   `json:"callGasLimit"`
		CallData             []string `json:"callData"`
		BlockTime            string   `json:"blockTime"`
		BlockNumber          string   `json:"blockNumber"`
		Beneficiary          string   `json:"beneficiary"`
		BaseFeePerGas        string   `json:"baseFeePerGas"`
		ActualGasUsed        string   `json:"actualGasUsed"`
		ActualGasCost        string   `json:"actualGasCost"`
		EntryPoint           string   `json:"entryPoint"`
		Erc20Transfers       []any    `json:"erc20Transfers"`
		Erc721Transfers      []any    `json:"erc721Transfers"`
		Value                []string `json:"value"`
		PreDecodedCallData   string   `json:"preDecodedCallData"`
	} `json:"userOps"`
}

func FetchOpsFromJiffyscan(transactionHash string) (*JiffyscanResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.jiffyscan.xyz/v0/getUserOp?hash=%s", transactionHash), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Host", "api.jiffyscan.xyz")
	req.Header.Add("x-api-key", "gFQghtJC6F734nPaUYK8M3ggf9TOpojkbNTH9gR5")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Origin", "https://www.jiffyscan.xyz")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15")
	req.Header.Add("Referer", "https://www.jiffyscan.xyz/")
	req.Header.Add("Accept", "*/*")

	if err := req.ParseForm(); err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response JiffyscanResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, err
}

func GetOnchainContractAddressFromDeploymentHash(aaHash string) (*string, error) {
	resp, err := FetchOpsFromJiffyscan(aaHash)
	if err != nil {
		return nil, err
	}

	var bloomHash string
	for _, bloom := range resp.UserOps {
		bloomHash = bloom.TransactionHash
	}

	contractAddress, err := addresswatcher.GetContractAddressFromParentHash(bloomHash)
	if err != nil {
		return nil, err
	}

	return contractAddress, nil
}

func StoreHashForDeployment(authDetails *internal.AuthDetails, input *model.DeploymentInfo) (*bool, error) {
	drop, err := engine.GetDropByID(input.DropID)
	if err != nil {
		return nil, errors.New("drop not found")
	}

	drop.AAWalletDeploymentHash = &input.DeploymentHash
	log.Info().Msgf("deployment info: %v", input)
	if input.ContractAddress == nil {
		// Introduce an artificial delay for before fethcing the actual contract address
		time.Sleep(time.Second * 3)

		contractAdddress, err := GetOnchainContractAddressFromDeploymentHash(input.DeploymentHash)
		if err != nil {
			log.Err(err)
		}

		drop.AAContractAddress = contractAdddress
	} else {
		drop.AAContractAddress = input.ContractAddress
	}

	err = engine.SaveModel(drop)
	if err != nil {
		return nil, err
	}

	t := true
	return &t, nil
}
