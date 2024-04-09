package neynar

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/lucidconnect/inverse/drops"
	"github.com/stretchr/testify/assert"
)

func TestNeynarClient_FetchFarcasterUserByUsername(t *testing.T) {
	godotenv.Load("../../.env.test.local")
	apiKeyOpt := WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := NewNeynarClient(apiKeyOpt)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	fid, err := neynarClient.FetchFarcasterUserByUsername("tezza")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !assert.Equal(t, int32(2037), fid) {
		t.FailNow()
	}
}

func TestNeynarClient_FetchFarcasterUserByEthAddress(t *testing.T) {
	godotenv.Load("../../.env.test.local")
	apiKeyOpt := WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := NewNeynarClient(apiKeyOpt)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	fid, err := neynarClient.FetchFarcasterUserFidByEthAddress("0xCcb9F5fAf66f15684A6154785f6aE524db6132E5")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !assert.Equal(t, int32(2037), fid) {
		t.FailNow()
	}
}

func TestNeynaClient_ValidateFarcasterChannelFollowers(t *testing.T) {
	godotenv.Load("../../.env.test")
	apiKeyOpt := WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := NewNeynarClient(apiKeyOpt)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	criteria := drops.FarcasterCriteria{
		ChannelID: "goat,higher",
	}
	valid := neynarClient.validateFarcasterChannelFollowerCriteria(2308, criteria)
	if !valid {
		t.FailNow()
	}
}
