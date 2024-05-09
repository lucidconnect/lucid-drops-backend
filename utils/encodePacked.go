package utils

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

// 🧑‍🎓 Scholared from
// https://gist.github.com/miguelmota/bc4304bb21a8f4cc0a37a0f9347b8bbb
// keccak256(abi.encodePacked( address,amount,message,nonce ))
// abi.EncodePacked 👇 wrappers
func EncodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func EncodeBytesString(v string) []byte {
	decoded, err := hex.DecodeString(v)
	if err != nil {
		log.Println("Failed for ", v)
		panic(err)
	}
	return decoded
}

func EncodeAddress(v string) []byte {
	decoded, err := hex.DecodeString(v[2:])
	if err != nil {
		log.Println("Failed for ", v)
		panic(err)
	}
	return decoded
}

func EncodeUint256(v string) []byte {
	fmt.Println(v)
	bn := new(big.Int)
	bn.SetString(v, 10)
	return math.U256Bytes(bn)
}
