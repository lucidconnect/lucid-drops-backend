package route

import (
	"fmt"
	"log"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/ledger"
	"inverse.so/models"
)

func CreditValidStripeWebhook(paymentIntent stripe.PaymentIntent, req *stripe.EventRequest) error {

	stripeWebhookBody := &models.StripeWebhooks{
		RequestID:      req.ID,
		IdempotencyKey: req.IdempotencyKey,
	}

	// get userID
	userID, ok := getUserIDFromCustomerID(paymentIntent.Customer.ID)
	if !ok || userID == "" {
		customer, err := getStripeCustomer(paymentIntent.Customer.ID)
		if err != nil {
			stripeWebhookBody.ErrorMetaData = fmt.Sprintf("customer not found from stripe %v\n", err)
			go engine.SaveModel(stripeWebhookBody)
			return err
		}

		// make sure customer is valid
		userID, ok = customer.Metadata["creatorId"]
		if !ok {
			stripeWebhookBody.ErrorMetaData = "creatorId not found in stripe webhook customer metadata"
			go engine.SaveModel(stripeWebhookBody)
			return fmt.Errorf("creatorId not found in customer metadata")
		}
	}

	//ledger instruction
	l := ledger.New(dbutils.DB)
	instruction := ledger.TransferInstruction{
		UserID: uuid.FromStringOrNil(userID),
		Amount: paymentIntent.Amount,
		Side:   ledger.Credit,
		TxRef:  paymentIntent.ID,
	}

	tx := dbutils.DB.Begin()
	err := l.Transfer(tx, instruction)
	if err != nil {
		stripeWebhookBody.ErrorMetaData = fmt.Sprintf("Ledger transfer error %v\n", err)
		go engine.SaveModel(stripeWebhookBody)
		tx.Rollback()
		return err
	}

	//TODO: create and update customer transactions in db
	customerID := paymentIntent.Customer.ID
	stripeWebhookBody.Processed = true
	go updateFirstPaymentAndCustomerIDStatus(userID, customerID)
	go engine.SaveModel(stripeWebhookBody)
	tx.Commit()
	return nil
}

func getUserIDFromCustomerID(customerID string) (string, bool) {

	var user models.Creator
	err := dbutils.DB.Model(&models.Creator{}).Where("stripe_customer_id", customerID).First(user).Error
	if err != nil {
		return "", false
	}

	return user.ID.String(), true
}

func getStripeCustomer(id string) (*stripe.Customer, error) {

	sc := &client.API{}
	sc.Init(os.Getenv("STRIPE_SECRET_KEY"), nil)
	customer, err := sc.Customers.Get(id, nil)
	if err != nil {
		log.Printf("Error getting customer %v\n", err)
		return nil, err
	}

	go persistStripeCustommerID(customer.Metadata["creatorId"], customer.ID)
	return customer, nil
}

func persistStripeCustommerID(userID, customerID string) error {

	var creator models.Creator
	err := dbutils.DB.Where("id = ?", userID).First(&creator).Error
	if err != nil {
		return err
	}

	creator.StripeCustomerID = &customerID
	err = dbutils.DB.Save(&creator).Error
	if err != nil {
		return err
	}

	return nil
}

func updateFirstPaymentAndCustomerIDStatus(userID, customerID string) error {

	var creator models.Creator
	err := dbutils.DB.Where("id = ?", userID).First(&creator).Error
	if err != nil {
		return err
	}

	creator.FirstPayment = true
	creator.StripeCustomerID = &customerID
	err = dbutils.DB.Save(&creator).Error
	if err != nil {
		return err
	}

	return nil
}
