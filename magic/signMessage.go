package magic

import (
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// TODO : Make this configurable
	samplePrivateKeyBytes = "0xafe2b745edf3ef97ffe9fb2fdf54253d1c3cb53f25ce3038a8e5f8a7b6ab367a"
)

func SecretlySignThisMessage(messageThatNeedsSigning string) (signature string) {
	privateKeyBytes := samplePrivateKeyBytes

	privateKey, err := crypto.HexToECDSA(privateKeyBytes[2:])
	if err != nil {
		log.Fatal("failed to parse PRIVATE_KEY", err.Error())
	}

	shaOfTheMessage := crypto.Keccak256([]byte(messageThatNeedsSigning))
	if err != nil {
		panic(err)
	}

	sig, err := crypto.Sign(shaOfTheMessage, privateKey)
	if err != nil {
		log.Printf("Sign error: %s", err)
	}

	// 💀 Transform yellow paper V from 27/28 to 0/1
	// vitalik did some stuff on the yellow paper thats why we need this
	sig[crypto.RecoveryIDOffset] += 27

	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	log.Println("Message      : 0x" + hex.EncodeToString(shaOfTheMessage))
	log.Println("Signature    : 0x" + hex.EncodeToString(sig))
	log.Println("🔐Signer  =    ", publicAddress)

	return hex.EncodeToString(sig)
}
