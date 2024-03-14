package whitelist

import (
	"errors"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/lucidNft"
	"github.com/lucidconnect/inverse/magic"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/utils"
	"github.com/rs/zerolog/log"
)

func GenerateSignatureForClaim(input *model.GenerateClaimSignatureInput, embeddedWalletAddress string) (*model.MintAuthorizationResponse, error) {

	now := time.Now()
	mintPass, err := engine.GetMintPassById(input.OtpRequestID)
	if err != nil {
		return nil, errors.New("mint pass not found")
	}

	if embeddedWalletAddress == "" {
		return nil, errors.New("embedded wallet address is required")
	}

	if strings.Contains(input.ClaimingAddress, ".eth") {
		resolvedAddress, err := utils.ResolveENSName(input.ClaimingAddress)
		if err != nil {
			return nil, err
		}

		input.ClaimingAddress = *resolvedAddress
	}

	if mintPass.UsedAt != nil {
		return nil, errors.New("mint pass has already been used")
	}

	var passes int64
	err = dbutils.DB.Model(&models.MintPass{}).Where("item_id = ? AND minter_address = ? AND used_at <> NULL", mintPass.DropID, input.ClaimingAddress).Count(&passes).Error
	if err == nil {
		if passes != 0 {
			return nil, errors.New("more than one mint pass found for this minter address")
		}
	}

	if IsThisAValidEthAddress(input.ClaimingAddress) {
		return nil, errors.New("the address in not a valid Ethereum address")
	}

	var addressClaiim models.WalletAddressClaim
	err = dbutils.DB.Where("item_id = ? AND wallet_address = ?", mintPass.DropID, input.ClaimingAddress).First(&addressClaiim).Error
	if err == nil {
		addressClaiim.EmbeddedWalletAddress = embeddedWalletAddress
		addressClaiim.SentOutAt = &now
		addressClaimError := engine.SaveModel(&addressClaiim)
		if addressClaimError != nil {
			log.Info().Msgf("🚨 Address Claim Model failed to updated in DB %+v", addressClaiim)
			return nil, errors.New("an error when verifying the Claim")
		}
	}

	tx := dbutils.DB.Begin()
	// userID, err := engine.GetCCreatorIDFromWalletAddress(embeddedWalletAddress)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, errors.New("claimer not found")
	// }

	// item, err := engine.GetItemByID(mintPass.DropID)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, errors.New("item not found")
	// }

	// if item.ClaimFee > 0 {
	// 	err = chargeClaimFee(*userID, item, tx)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return nil, err
	// 	}
	// }

	mintPass.MinterAddress = input.ClaimingAddress
	mintPass.UsedAt = &now
	mintPassSaveError := engine.SaveModel(&mintPass)
	if mintPassSaveError != nil {
		tx.Rollback()
		log.Info().Msgf("🚨 Mint Pass Model failed to updated in DB %+v", mintPass)
		return nil, errors.New("an error when verifying the Claim")
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// TODO add back signature flow
	// (claim_address+contract_address+chain_id+amount+number_of_mints)
	contractAddress := common.HexToAddress(mintPass.DropContractAddress)
	chainId := getChainId(*mintPass.BlockchainNetwork)

	caller := getEthBackend(os.Getenv("RPC_PROVIDER"))
	lucidNftCaller, err := lucidNft.NewLucidNftCaller(contractAddress, caller)
	if err != nil {
		log.Err(err).Caller().Send()
		return nil, err
	}

	mints, err := lucidNftCaller.GetMints(nil, common.Big1)
	if err != nil {
		log.Err(err).Caller().Send()
		return nil, err
	}

	dataInPackedFormat := utils.EncodePacked(
		// utils.EncodeAddress("0x14723A09ACff6D2A60DcdF7aA4AFf308FDDC160"),
		utils.EncodeAddress(input.ClaimingAddress), // Some Addresss
		utils.EncodeAddress(contractAddress.Hex()),
		chainId.Bytes(),
		common.Big1.Bytes(),
		mints.Bytes(),
	)

	rawData := hexutil.Encode(dataInPackedFormat)

	keccakOfTheMessageInBytes := crypto.Keccak256(dataInPackedFormat)

	signature := magic.SecretlySignThisMessage("\x19Ethereum Signed Message:\n32" + string(keccakOfTheMessageInBytes))

	return &model.MintAuthorizationResponse{
		Amount:               "1",
		TokenID:              "1",
		Nonce:                mints.String(),
		Chain:                int(chainId.Int64()),
		PackedData:           rawData,
		MintingAbi:           "['function mint(address _to) public']",
		MintingSignature:     signature,
		SmartContractAddress: contractAddress.Hex(),
	}, nil
}

func getEthBackend(rpc string) *ethclient.Client {
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		log.Err(err).Msg("Failed to connect to the Ethereum client")
	}
	return conn
}

func getChainId(network model.BlockchainNetwork) *big.Int {
	var chain *big.Int
	switch network {
	case model.BlockchainNetworkBase:
		if ok, _ := utils.IsProduction(); ok {
			chain = big.NewInt(8453)
		} else {
			chain = big.NewInt(84532)
		}
	default:
		return nil
	}
	return chain
}

func GenerateSignatureForFarcasterClaim(input *model.GenerateClaimSignatureInput) (*model.MintAuthorizationResponse, error) {

	now := time.Now()
	mintPass, err := engine.GetMintPassById(input.OtpRequestID)
	if err != nil {
		return nil, errors.New("mint pass not found")
	}

	// if strings.Contains(input.ClaimingAddress, ".eth") {
	// 	resolvedAddress, err := utils.ResolveENSName(input.ClaimingAddress)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	input.ClaimingAddress = *resolvedAddress
	// }

	if mintPass.UsedAt != nil {
		return nil, errors.New("mint pass has already been used")
	}

	var passes int64
	err = dbutils.DB.Model(&models.MintPass{}).Where("item_id = ? AND minter_address = ? AND used_at <> NULL", mintPass.DropID, input.ClaimingAddress).Count(&passes).Error
	if err == nil {
		if passes != 0 {
			return nil, errors.New("more than one mint pass found for this minter address")
		}
	}

	if IsThisAValidEthAddress(input.ClaimingAddress) {
		return nil, errors.New("the address in not a valid Ethereum address")
	}

	// var addressClaiim models.WalletAddressClaim
	// err = dbutils.DB.Where("item_id = ? AND wallet_address = ?", mintPass.ItemId, input.ClaimingAddress).First(&addressClaiim).Error
	// if err == nil {
	// 	addressClaiim.EmbeddedWalletAddress = embeddedWalletAddress
	// 	addressClaiim.SentOutAt = &now
	// 	addressClaimError := engine.SaveModel(&addressClaiim)
	// 	if addressClaimError != nil {
	// 		log.Info().Msgf("🚨 Address Claim Model failed to updated in DB %+v", addressClaiim)
	// 		return nil, errors.New("an error when verifying the Claim")
	// 	}
	// }

	tx := dbutils.DB.Begin()

	mintPass.MinterAddress = input.ClaimingAddress
	mintPass.UsedAt = &now
	mintPassSaveError := engine.SaveModel(&mintPass)
	if mintPassSaveError != nil {
		tx.Rollback()
		log.Info().Msgf("🚨 Mint Pass Model failed to updated in DB %+v", mintPass)
		return nil, errors.New("an error when verifying the Claim")
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// TODO add back signature flow
	// (claim_address+contract_address+chain_id+amount+number_of_mints)
	contractAddress := common.HexToAddress(mintPass.DropContractAddress)
	chainId := getChainId(*mintPass.BlockchainNetwork)

	caller := getEthBackend(os.Getenv("RPC_PROVIDER"))
	lucidNftCaller, err := lucidNft.NewLucidNftCaller(contractAddress, caller)
	if err != nil {
		log.Err(err).Caller().Send()
		return nil, err
	}

	mints, err := lucidNftCaller.GetMints(nil, common.Big1)
	if err != nil {
		log.Err(err).Caller().Send()
		return nil, err
	}

	dataInPackedFormat := utils.EncodePacked(
		utils.EncodeAddress(input.ClaimingAddress), // Some Addresss
		utils.EncodeAddress(contractAddress.Hex()),
		utils.EncodeUint256(chainId.String()),
		utils.EncodeUint256("1"),
		utils.EncodeUint256(mints.String()),
	)

	rawData := hexutil.Encode(dataInPackedFormat)

	keccakOfTheMessageInBytes := crypto.Keccak256(dataInPackedFormat)

	signature := magic.SecretlySignThisMessage("\x19Ethereum Signed Message:\n32" + string(keccakOfTheMessageInBytes))

	return &model.MintAuthorizationResponse{
		Amount:               "1",
		TokenID:              "1",
		Nonce:                mints.String(),
		Chain:                int(chainId.Int64()),
		PackedData:           rawData,
		MintingAbi:           "['function mint(address _to) public']",
		MintingSignature:     signature,
		SmartContractAddress: contractAddress.Hex(),
	}, nil
}
