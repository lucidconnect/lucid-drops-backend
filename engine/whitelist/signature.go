package whitelist

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/magic"
	"inverse.so/utils"
)

func GenerateSignatureForClaim(input *model.GenerateClaimSignatureInput) (*model.MintAuthorizationResponse, error) {
	mintPass, err := engine.GetMintPassById(input.OtpRequestID)
	if err != nil {
		return nil, errors.New("mint pass not found")
	}

	if mintPass.UsedAt != nil {
		return nil, errors.New("mint pass has already been used")
	}

	if IsThisAValidEthAddress(input.ClaimingAddress) {
		return nil, errors.New("the passed in address in not a valid Ethereum address")
	}

	mintPass.MinterAddress = input.ClaimingAddress

	mintPassSaveError := engine.SaveModel(mintPass)
	if mintPassSaveError != nil {
		log.Info().Msgf("🚨 Mint Pass Model failed to updated in DB %+v", mintPass)
		return nil, errors.New("an error when verifying the Claim")
	}

	go func() {
		inverseAAServerURL := utils.UseEnvOrDefault("INVERSE_AA_SERVER", "https://inverse-aa.onrender.com")

		client := &http.Client{}

		itemData, err := json.Marshal(map[string]interface{}{
			"receiptientAddresses": []string{input.ClaimingAddress},
			"items":                []int64{mintPass.ItemIdOnContract},
			"contractAddress":      mintPass.CollectionContractAddress,
			"Network":              mintPass.BlockchainNetwork,
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := http.NewRequest(http.MethodPost, inverseAAServerURL+"/sendnfts", bytes.NewBuffer(itemData))
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer res.Body.Close()
	}()

	// TODO add back signature flow
	dataInPackedFormat := utils.EncodePacked(
		// utils.EncodeAddress("0x14723A09ACff6D2A60DcdF7aA4AFf308FDDC160"),
		utils.EncodeAddress(input.ClaimingAddress), // Some Addresss
		utils.EncodeUint256("123"),
		utils.EncodeBytesString(hex.EncodeToString([]byte("coffee and donuts"))),
		utils.EncodeUint256("1"),
	)

	rawData := hexutil.Encode(dataInPackedFormat)

	keccakOfTheMessageInBytes := crypto.Keccak256(dataInPackedFormat)

	signature := magic.SecretlySignThisMessage("\x19Ethereum Signed Message:\n32" + string(keccakOfTheMessageInBytes))

	return &model.MintAuthorizationResponse{
		PackedData:           rawData,
		MintingAbi:           "['function mint(address _to) public']",
		MintingSignature:     signature,
		SmartContractAddress: "0x34bE7f35132E97915633BC1fc020364EA5134863",
	}, nil
}
