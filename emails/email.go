package emails

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/mailgun/mailgun-go"
	"github.com/rs/zerolog/log"
)

type EmailData struct {
	Subject     string
	ContentData interface{}
	EmailTo     string
	EmailFrom   string
	Template    string
}

type EmailResponse struct {
	MessageID string
	Response  string
}

// this function parses variables into the respective templates
func processTemplate(templatePath string, dynamicData interface{}) (string, error) {
	var temp bytes.Buffer
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	tmpl.Execute(&temp, dynamicData)
	return temp.String(), nil
}

func SendVerificationEmail(emailAddress, fromEmailAdress, otp, action, expiry string) error {

	templatePath := "emails/templates/emailverification.html"
	dynamicData := map[string]interface{}{
		"otp":    otp,
		"action": action,
		"expiry": expiry,
	}

	body, err := processTemplate(templatePath, dynamicData)
	if err != nil {
		return err
	}

	data := EmailData{
		Subject:     "Verify your email address",
		ContentData: body,
		EmailTo:     emailAddress,
		EmailFrom:   fromEmailAdress,
		Template:    templatePath,
	}
	_, err = SendEmail(data)
	if err != nil {
		return err
	}

	return nil
}

func SendEmail(object EmailData) (*EmailResponse, error) {
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")

	var err error
	mg := mailgun.NewMailgun(domain, apiKey)

	sender := object.EmailFrom
	from := fmt.Sprintf("Inverse <%s>", sender)
	subject := object.Subject
	recipient := object.EmailTo
	body := fmt.Sprintf("%v", object.ContentData)

	message := mg.NewMessage(from, subject, body, recipient)
	message.SetHtml(body)

	response, messageId, err := mg.Send(message)
	if err != nil {
		log.Err(err).Msg("Error sending email")
		return nil, err
	}

	return &EmailResponse{
		MessageID: messageId,
		Response:  response,
	}, nil
}

func SendClaimNudgeEmail(emailAddress, fromEmailAdress, itemName, claimLink, creatorUsername string) error {
	templatePath := "emails/templates/inverse-email-criteria.html"
	dynamicData := map[string]interface{}{
		"itemName":        itemName,
		"creatorUsername": creatorUsername,
		"claimLink":       claimLink,
	}

	body, err := processTemplate(templatePath, dynamicData)
	if err != nil {
		return err
	}

	data := EmailData{
		Subject:     "You've got an item",
		ContentData: body,
		EmailTo:     emailAddress,
		EmailFrom:   fromEmailAdress,
		Template:    templatePath,
	}
	_, err = SendEmail(data)
	if err != nil {
		return err
	}

	return nil
}
