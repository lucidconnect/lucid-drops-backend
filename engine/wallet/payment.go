package wallet

import (
	"github.com/lucidconnect/inverse/engine/onboarding"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/services"
)

func GenerateCreatorPaymentIntentSecret(authDetails *internal.AuthDetails, amount int64) (*string, error) {

	customerID, err := onboarding.FindOrCreateStripeCustomerID(authDetails)
	if err != nil {
		return nil, err
	}

	paymentIntent, err := services.CreateStripePaymentIntent(amount, "usd", *customerID)
	if err != nil {
		return nil, err
	}

	return &paymentIntent.ClientSecret, nil
}
