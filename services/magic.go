package services

import (
	"time"

	"github.com/lucidconnect/inverse/utils"
	"github.com/magiclabs/magic-admin-go"
	"github.com/magiclabs/magic-admin-go/client"
	"github.com/magiclabs/magic-admin-go/token"
)

func GetMagicClient() *client.API {
	cl := magic.NewClientWithRetry(5, time.Second, 10*time.Second)

	c, er := client.New(utils.UseEnvOrDefault("MAGIC_SECRET_KEY", "_default string"), cl)
	if er != nil {
		return nil
	}
	return c
}

func GenerateMagicJWT(didToken string) (*token.Token, error) {

	decoded, err := token.NewToken(didToken)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func GetMagicAddress(resolvedToken *token.Token) (*string, error) {
	err := resolvedToken.Validate("")
	if err != nil {
		return nil, err
	}

	publicAddress, err := resolvedToken.GetPublicAddress()
	if err != nil {
		return nil, err
	}

	return &publicAddress, nil
}
