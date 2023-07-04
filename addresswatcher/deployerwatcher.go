package addresswatcher

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"inverse.so/engine"
	"inverse.so/utils"
)

type LogContractCreation struct {
	NFTAddress common.Address
}

func SubscribeToInverseContractDeployments() {
	rpcProvider := utils.UseEnvOrDefault("RPC_PROVIDER", "wss://polygon-mumbai.g.alchemy.com/v2/SjkYprQJA0Dp-5l5XpafAmqj_uoJ_A5G")
	inveseNFTFactoryAddress := utils.UseEnvOrDefault("INVERSE_FACTORY_ADDRESS", "0x021406A44658CAbcBc5540Ec2045123E5FDb0ca8")

	client, err := ethclient.Dial(rpcProvider)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	contractAddress := common.HexToAddress(inveseNFTFactoryAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	log.Info().Msgf("🪼 Started watcher for (%s) Contract", inveseNFTFactoryAddress)

	contractDeploymentHash := crypto.Keccak256Hash([]byte("TokenDeployed(address)"))

	for {
		select {
		case err := <-sub.Err():
			log.Error().Msg(err.Error())
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
