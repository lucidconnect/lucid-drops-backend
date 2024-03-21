package neynar

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
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
