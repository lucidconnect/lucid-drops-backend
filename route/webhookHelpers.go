package route

import (
	"fmt"
	"log"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
	"inverse.so/dbutils"
	"inverse.so/ledger"
	"inverse.so/models"
)

func CreditValidStripeWebhook(paymentIntent stripe.PaymentIntent) error {

	// get customer from stripe
	//TODO: get customer from db cache
	customer, err := getStripeCustomer(paymentIntent.Customer.ID)
	if err != nil {
		return err
	}

	// make sure customer is valid
	customerID, ok := customer.Metadata["creatorId"]
	if !ok {
		return fmt.Errorf("creatorId not found in customer metadata")
	}

	//ledger instruction
	l := ledger.New(dbutils.DB)
	instruction := ledger.TransferInstruction{
		UserID: uuid.FromStringOrNil(customerID),
		Amount: paymentIntent.Amount,
		Side:   ledger.Credit,
		TxRef:  paymentIntent.ID,
	}

	tx := dbutils.DB.Begin()
	err = l.Transfer(tx, instruction)
	if err != nil {
		tx.Rollback()
		return err
	}

	//TODO: create and update customer transactions in db
	go updateFirstPaymentStatus(customerID)
	tx.Commit()
	return nil
}

func getStripeCustomer(id string) (*stripe.Customer, error) {

	sc := &client.API{}
	sc.Init(os.Getenv("STRIPE_SECRET_KEY"), nil)
	customer, err := sc.Customers.Get(id, nil)
	if err != nil {
		log.Printf("Error getting customer %v\n", err)
		return nil, err
	}

	return customer, nil
}

func updateFirstPaymentStatus(customerID string) error {

	var creator models.Creator
	err := dbutils.DB.Where("id = ?", customerID).First(&creator).Error
	if err != nil {
		return err
	}

	creator.FirstPayment = true
	err = dbutils.DB.Save(&creator).Error
	if err != nil {
		return err
	}

	return nil
}
