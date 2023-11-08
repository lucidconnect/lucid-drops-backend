package mobile

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
)

type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

func GenerateRandomEthAddress() (*KeyPair, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	newKeyPair := new(KeyPair)

	newKeyPair.PrivateKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	newKeyPair.PublicKey = address

	return newKeyPair, nil
}

func GenerateMobileWalletConfigs(authDetails *internal.AuthDetails) (*model.MobileWalletConfig, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new collection")
	}

	altSigner, err := engine.GetAltSignerByCreatorID(creator.ID.String())
	if err != nil {
		return nil, err
	}

	if altSigner.AltPrivateKey != "" {

		return &model.MobileWalletConfig{
			PublicKey:  altSigner.AltPublicKey,
			PrivateKey: altSigner.AltPrivateKey,
			AaWallet:   altSigner.WalletAddress,
		}, err
	}

	accountDetails, generationErr := GenerateRandomEthAddress()
	if generationErr != nil {
		return nil, generationErr
	}

	altSigner.AltPublicKey = accountDetails.PublicKey
	altSigner.AltPrivateKey = accountDetails.PrivateKey

	err = engine.SaveModel(nil, altSigner)
	if err != nil {
		return nil, err
	}

	return &model.MobileWalletConfig{
		PublicKey:  accountDetails.PublicKey,
		PrivateKey: accountDetails.PrivateKey,
		AaWallet:   altSigner.WalletAddress,
	}, nil
}
