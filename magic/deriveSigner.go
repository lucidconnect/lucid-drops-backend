package magic

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetMeTheSignerOfThisMessage(originalPackedMessage string, signatureOfTheMessageShaInHex string) (*common.Address, error) {
	sigInHex := hexutil.MustDecode(signatureOfTheMessageShaInHex)

	shaOfTheMessage := crypto.Keccak256([]byte(originalPackedMessage))

	// 💀 Transform yellow paper V from 27/28 to 0/1
	// vitalik did some stuff on the yellow paper thats why we need this
	if sigInHex[crypto.RecoveryIDOffset] == 27 || sigInHex[crypto.RecoveryIDOffset] == 28 {
		sigInHex[crypto.RecoveryIDOffset] -= 27
	}

	recoveredPub, err := crypto.Ecrecover(shaOfTheMessage, sigInHex)
	if err != nil {
		log.Printf("ECRecover error: %s", err)
		return nil, err
	}

	pubKey, err := crypto.UnmarshalPubkey(recoveredPub)
	if err != nil {
		log.Printf("Unmarshall error: %s", err)
		return nil, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return &recoveredAddr, nil
}
