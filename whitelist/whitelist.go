package whitelist

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lucidconnect/inverse/drops"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/lucidNft"
	"github.com/lucidconnect/inverse/magic"
	"github.com/lucidconnect/inverse/utils"
	"github.com/rs/zerolog/log"
)

func GenerateSignatureForClaim(mintPass drops.MintPass) (*model.MintAuthorizationResponse, error) {
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
		utils.EncodeAddress(mintPass.MinterAddress), // Some Addresss
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

func isThisAValidEthAddress(maybeAddress string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	if len(maybeAddress) != 43 {
		return false
	}

	return re.MatchString(maybeAddress)
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

func MintNft(mintArgs model.MintAuthorizationResponse, walletAddress string) (string, error) {
	claimer := common.HexToAddress(walletAddress)
	contract := common.HexToAddress(mintArgs.SmartContractAddress)

	amount, _ := new(big.Int).SetString(mintArgs.Amount, 10)
	tokenId, _ := new(big.Int).SetString(mintArgs.TokenID, 10)
	nonce, _ := new(big.Int).SetString(mintArgs.Nonce, 10)
	// amount.SetString(signatureResponse.Amount, 10)
	// tokenId.SetString(signatureResponse.TokenID, 10)
	// nonce.SetString(signatureResponse.Nonce, 10)
	signature, err := hex.DecodeString(mintArgs.MintingSignature)
	if err != nil {
		err = fmt.Errorf("decoding signature failed with error %v", err)
		log.Info().Msgf("signature %v", mintArgs.MintingSignature)
		log.Err(err).Send()
		return "", err
	}

	return mint(contract, claimer, amount, tokenId, nonce, signature, int64(mintArgs.Chain))
}

func mint(contractAddress, claimAddress common.Address, amount, tokenId, nonce *big.Int, signature []byte, chain int64) (string, error) {
	rpc := os.Getenv("RPC_PROVIDER")
	privateKey := os.Getenv("PRIVATE_KEY")
	backend := getEthBackend(rpc)
	lucidNftCaller, err := lucidNft.NewLucidNft(contractAddress, backend)
	if err != nil {
		log.Err(err).Send()
		return "", err
	}

	transactOpt, err := getTransactor(privateKey, chain)
	if err != nil {
		log.Err(err).Send()
		return "", err
	}
	tx, err := lucidNftCaller.Mint(transactOpt, claimAddress, amount, tokenId, nonce, signature)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func getTransactor(privateKey string, chain int64) (*bind.TransactOpts, error) {
	key, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		return nil, err
	}

	chainId := big.NewInt(chain)

	transactionOpts, err := bind.NewKeyedTransactorWithChainID(key, chainId)
	if err != nil {
		return nil, err
	}
	return transactionOpts, nil
}
