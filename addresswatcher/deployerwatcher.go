package addresswatcher

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/utils"
)

var globalRPCClient *ethclient.Client

type LogContractCreation struct {
	NFTAddress common.Address
}

func SubscribeToInverseContractDeployments() {
	rpcProvider := utils.UseEnvOrDefault("RPC_PROVIDER", "wss://polygon-mainnet.g.alchemy.com/v2/JKEn31gUjMUuyOjgCGpTNS-fJnGp-mYF")
	inveseNFTFactoryAddress := utils.UseEnvOrDefault("INVERSE_FACTORY_ADDRESS", "0x021406A44658CAbcBc5540Ec2045123E5FDb0ca8")

attemptReconnect:
	var err error
	globalRPCClient, err = ethclient.Dial(rpcProvider)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	contractAddress := common.HexToAddress(inveseNFTFactoryAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := globalRPCClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Error().Msg(err.Error())
		goto attemptReconnect
	}

	log.Info().Msgf("🪼 Started watcher for (%s) Contract", inveseNFTFactoryAddress)

	contractDeploymentHash := crypto.Keccak256Hash([]byte("TokenDeployed(address)"))

	for {
		select {
		case err := <-sub.Err():
			log.Error().Msg(err.Error())
			goto attemptReconnect
		case vLog := <-logs:
			eventType := vLog.Topics[0].Hex()
			if eventType != contractDeploymentHash.Hex() {
				log.Info().Msg("😏 A new Inverse NFT has been deployed")
				continue
			}

			contractAbi, err := abi.JSON(strings.NewReader(string(AddresswatcherABI)))
			if err != nil {
				log.Error().Msg(err.Error())
			}

			var creationEvent LogContractCreation
			err = contractAbi.UnpackIntoInterface(&creationEvent, "TokenDeployed", vLog.Data)
			if err != nil {
				log.Error().Msg(err.Error())
			}

			deployedContractAddress := creationEvent.NFTAddress

			engine.AttachContractAddressForCreationHash(vLog.TxHash.String(), deployedContractAddress.String())
			log.Info().Msgf("🔖 (%s) Contract Deployed %s", vLog.TxHash, deployedContractAddress)
		}
	}
}
func GetContractAddressFromParentHash(onchainHash string) (*string, error) {
	if globalRPCClient == nil {
		rpcProvider := utils.UseEnvOrDefault("RPC_PROVIDER", "wss://polygon-mainnet.g.alchemy.com/v2/98DqwUXmmWt8TZc7sbDSwHGdBAY0PeuF")

		var err error
		globalRPCClient, err = ethclient.Dial(rpcProvider)
		if err != nil {
			return nil, err
		}
	}

	// Replace this with your transaction hash
	txHash := common.HexToHash(onchainHash)

	// Get the transaction receipt
	receipt, err := globalRPCClient.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		fmt.Println("Error fetching transaction receipt:", err)
		return nil, err
	}

	contractDeploymentHash := crypto.Keccak256Hash([]byte("TokenDeployed(address)"))

	if len(receipt.Logs) == 0 {
		return nil, errors.New("transaction has no logs")
	}

	for _, vLog := range receipt.Logs {
		eventType := vLog.Topics[0].Hex()
		if eventType != contractDeploymentHash.Hex() {
			continue
		}
		contractAbi, err := abi.JSON(strings.NewReader(string(AddresswatcherABI)))
		if err != nil {
			log.Error().Msg(err.Error())
		}

		var creationEvent LogContractCreation
		err = contractAbi.UnpackIntoInterface(&creationEvent, "TokenDeployed", vLog.Data)
		if err != nil {
			log.Error().Msg(err.Error())
		}

		deployedContractAddress := creationEvent.NFTAddress.String()

		return &deployedContractAddress, nil
	}

	return nil, errors.New("transaction has no logs")
}
