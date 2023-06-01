package magic

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetMeTheapitypesOfThisMessage(originalPackedMessage string, signatureOfTheMessageShaInHex string) common.Address {
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
	}

	pubKey, err := crypto.UnmarshalPubkey(recoveredPub)
	if err != nil {
		log.Printf("Unmarshall error: %s", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	log.Printf("Message Signer has been recoverd and its %s", recoveredAddr)

	return recoveredAddr
}
