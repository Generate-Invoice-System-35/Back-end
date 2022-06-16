package usecase

import (
	"fmt"
	"log"

	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"gopkg.in/gomail.v2"
)

type serviceSendCustomer struct {
	c    config.Config
	repo adapter.AdapterSendCustomerRepository
}

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Fatture Generate Invoices <fattureinvoices35@gmail.com>"
const CONFIG_AUTH_EMAIL = "fattureinvoices35@gmail.com"
const CONFIG_AUTH_PASSWORD = "FattureInvoices123456789"

func (s *serviceSendCustomer) SendEmailService(msg model.SendCustomer) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", msg.To)
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", msg.Subject)
	mailer.SetBody("text/html", msg.Body)
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
	// SMTP
	// email := "fattureinvoices35@gmail.com"
	// from := "fattureinvoices35@gmail.com"
	// password := "FattureInvoices123456789"

	// toEmailAdress := msg.To
	// to := []string{toEmailAdress}

	// host := "smtp.mailtrap.io"
	// port := "25"
	// address := host + ":" + port

	// subject := "Subject: " + msg.Subject + "\n"
	// body := msg.Body
	// message := []byte(subject + body)

	// auth := smtp.PlainAuth("", email, password, host)
	// err := smtp.SendMail(address, auth, from, to, message)
	// if err != nil {
	// 	panic(err)
	// }

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
