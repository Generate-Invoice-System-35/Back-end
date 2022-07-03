package usecase

import (
	"fmt"
	"log"
	"net/smtp"
	"time"

	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type serviceSendCustomer struct {
	c    config.Config
	repo adapter.AdapterSendCustomerRepository
}

// const CONFIG_SMTP_HOST = "smtp.gmail.com"
// const CONFIG_SMTP_PORT = 587
// const CONFIG_SENDER_NAME = "Fatture Generate Invoices <fattureinvoices35@gmail.com>"
// const CONFIG_AUTH_EMAIL = "fattureinvoices35@gmail.com"
// const CONFIG_AUTH_PASSWORD = "FattureInvoices123456789"

func (s *serviceSendCustomer) SendEmailService(msg model.SendCustomer) error {
	server := "smtp-mail.outlook.com"
	port := 587
	user := "fattureinvoices35@outlook.com"
	from := user
	pass := "FattureInvoices123456789"
	dest := msg.To

	auth := LoginAuth(user, pass)

	to := []string{dest}

	message := []byte("From: " + from + "\n" +
		"To: " + dest + "\n" +
		"Subject: " + msg.Subject + "\n\n" +
		msg.Body)

	endpoint := fmt.Sprintf("%v:%v", server, port)
	err := smtp.SendMail(endpoint, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}

	msg.Created_At = time.Now()
	msg.Updated_At = time.Now()
	return s.repo.SendEmail(msg)
}

func (s *serviceSendCustomer) SendWhatsappService(msg model.SendCustomer) error {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:+62" + msg.To)
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody(msg.Body)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		log.Println("Message sent successfully!")
	}

	return s.repo.SendWhatsapp(msg)
}

func NewServiceSendCustomer(repo adapter.AdapterSendCustomerRepository, c config.Config) adapter.AdapterSendCustomerService {
	return &serviceSendCustomer{
		repo: repo,
		c:    c,
	}
}
