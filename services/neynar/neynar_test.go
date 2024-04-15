package neynar

import (
	"fmt"
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

func TestNeynarClient_ValidateFarcasterChannelFollowers(t *testing.T) {
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

func TestNeynarClient_ValidateFarcasterLikeCriteria(t *testing.T) {
	godotenv.Load("../../.env.test")
	apiKeyOpt := WithNeynarApiKey(os.Getenv("NEYNAR_API_KEY"))
	neynarClient, err := NewNeynarClient(apiKeyOpt)
	if !assert.NoError(t, err) {
		t.Logf("client initialization error %v", err)
		t.FailNow()
	}

	farcasterAddress := "0xb77ce6ec08b85dcc468b94cea7cc539a3bbf9510"

	userFid, err := neynarClient.FetchFarcasterUserFidByEthAddress(farcasterAddress)
	if !assert.NoError(t, err) {
		t.Logf("error fetching farcaster user %v", err)
		t.Fail()
	}

	if !assert.Equal(t, userFid, int32(2037)) {
		t.Logf("fid %v", userFid)
		t.FailNow()
	}

	criteria := drops.FarcasterCriteria{
		CastUrl: "https://warpcast.com/shay/0x63a6e387",
	}

	valid := neynarClient.validateFarcasterLikeCriteria(userFid, criteria)
	fmt.Println(valid)
	assert.Equal(t, valid, true)
	if !valid {
		t.Logf("user is eligible - %v", valid)
		t.FailNow()
	}
}
