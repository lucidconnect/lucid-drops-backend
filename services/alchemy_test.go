package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlchemyClient_GetOwnersForNft(t *testing.T) {
	var options []Option
	apiKeyOption := WithApiKey("cILOyabXlHIQfSz39zQ5zmwsVANoAn06")
	urlOption := WithUrl("https://base-sepolia.g.alchemy.com")
	options = append(options, apiKeyOption, urlOption)
	testAlchemyClient, err := NewAlchemyClient(options...)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	owners, err := testAlchemyClient.GetOwnersForNft("0x8305b07d1083efA2b493BC2304B23d23338b3cdA", "1")
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	fmt.Println("owners: ", owners)
}
