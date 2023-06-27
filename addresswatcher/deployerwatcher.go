package addresswatcher

import (
	"context"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"inverse.so/engine"
)

const (
	abacusFactoryAddress = "0x095654d842262944a0e2d93E1D8394C5d459255a"
)

type LogContractCreation struct {
	NFTAddress common.Address
}

func SubscribeToInverseContractDeployments() {
	client, err := ethclient.Dial("wss://polygon-mumbai.g.alchemy.com/v2/SjkYprQJA0Dp-5l5XpafAmqj_uoJ_A5G")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(abacusFactoryAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🪼 Started Watcher")

	contractDeploymentHash := crypto.Keccak256Hash([]byte("TokenDeployed(address)"))

	for {
		select {
		case err := <-sub.Err():
			log.Println(err)
		case vLog := <-logs:
			eventType := vLog.Topics[0].Hex()
			if eventType == contractDeploymentHash.Hex() {
				log.Println("😏 A new Inverse NFT has been deployed")
			} else {
				log.Println("😅Shit has hit the fan", eventType)
			}

			contractAbi, err := abi.JSON(strings.NewReader(string(AddresswatcherABI)))
			if err != nil {
				log.Fatal(err)
			}

			var creationEvent LogContractCreation
			err = contractAbi.UnpackIntoInterface(&creationEvent, "TokenDeployed", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			deployedContractAddress := creationEvent.NFTAddress

			engine.AttachContractAddressForCreationHash(vLog.TxHash.String(), deployedContractAddress.String())
			log.Printf("🔖 (%s) Contract Deployed %s", vLog.TxHash, deployedContractAddress)
		}
	}
}
