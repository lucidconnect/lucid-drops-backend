package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"inverse.so/mintwatcher"
	"inverse.so/utils"
)

var ContractAddress = common.HexToAddress("0x783D65396Ea7c260FF7D1ABbcc626c4D9c7A2FB2")

func TestMain(m *testing.M) {

	rpcProvider := "https://polygon.gateway.tenderly.co/38MxkoHCeZcY74yGFmpZ4D"
	client, err := ethclient.Dial(rpcProvider)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	utils.SetUpDefaultLogger()
	err = initializeMintWatcher(client)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	m.Run()
}

func initializeMintWatcher(client *ethclient.Client) error {

	var err error
	mintwatcher.LocalTestWatcher, err = mintwatcher.NewMintwatcher(ContractAddress, client)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

func initializeDynamicMintWatcher(address string, client *ethclient.Client) error {
	var err error
	mintwatcher.LocalTestWatcher, err = mintwatcher.NewMintwatcher(common.HexToAddress(address), client)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}
