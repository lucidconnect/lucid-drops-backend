package services

import (
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

func CreateStripeCustomerID(creatorID string) (*stripe.Customer, error) {

	params := &stripe.CustomerParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"creatorId": creatorID,
			},
		},
	}

	sc := &client.API{}
	sc.Init(os.Getenv("STRIPE_SECRET_KEY"), nil)
	customer, err := sc.Customers.New(params)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func CreateStripePaymentIntent(amount int64, currency string, customerID string) (*stripe.PaymentIntent, error) {

	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(amount),
		Currency:      stripe.String(currency),
		Customer:      stripe.String(customerID),
		PaymentMethod: stripe.String("card"),
	}

	sc := &client.API{}
	sc.Init(os.Getenv("STRIPE_SECRET_KEY"), nil)
	paymentIntent, err := sc.PaymentIntents.New(params)
	if err != nil {
		return nil, err
	}

	return paymentIntent, nil
}
