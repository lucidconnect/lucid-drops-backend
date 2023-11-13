package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v3"
)

func ResolveENSName(domain string) (*string, error) {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/b67c22367fde461d8916fd676604f16d")
	if err != nil {
		return nil, err
	}

	address, err := ens.Resolve(client, domain)
	if err != nil {
		return nil, err
	}

	addressHex := address.Hex()

	return &addressHex, nil
}
