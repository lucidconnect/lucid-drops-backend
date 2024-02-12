// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/rs/zerolog/log"
	"github.com/lucidconnect/inverse/mintwatcher"
)

func TestCorrectURI(t *testing.T) {
	actualURI := "https://inverse-prod.onrender.com/metadata/0x783D65396Ea7c260FF7D1ABbcc626c4D9c7A2FB2/094a8eae-c2b0-4ff7-b11e-37e9fd300036"
	log.Info().Msg("Testing correct Token URI")
	got, err := mintwatcher.LocalTestWatcher.Uri(&bind.CallOpts{}, big.NewInt(1))
	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, actualURI, got)
}

func TestInCorrectURI(t *testing.T) {
	actualURI := "https://inverse-prod.onrender.com/metadata/0x783D65396Ea7c260FF7D1ABbcc626c4D9c7A2FB2/6e37c8fe-f23d-45f9-b2d7-3444b65b40bc"
	fakeURI := "https://inverse-prod.onrender.com/metadata/0x783D65396Ea7c260FF7D1ABbcc626c4D9c7A2FB2/094a8eae-c2b0-4ff7-b11e-37e9fd300036"
	log.Info().Msg("Testing correct Token URI")
	got, err := mintwatcher.LocalTestWatcher.Uri(&bind.CallOpts{}, big.NewInt(2))
	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, actualURI, got)
	assert.NotEqual(t, fakeURI, got)
}

func TestOWner(t *testing.T) {
	owner := "0x98cD55827F2b6d7745586A62CFD2d6a5077472ed"
	contractAddress := "0x783D65396Ea7c260FF7D1ABbcc626c4D9c7A2FB2"
	got, err := mintwatcher.LocalTestWatcher.Owner(&bind.CallOpts{})
	log.Info().Msgf("Owner: %s", got.Hex())
	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, owner, got.Hex())
	assert.NotEqual(t, contractAddress, got.Hex())
}
