package usecase

import (
	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"
	"net/smtp"
)

type serviceSendCustomer struct {
	c    config.Config
	repo adapter.AdapterSendCustomerRepository
}

func (s *serviceSendCustomer) SendEmailService(msg model.SendCustomer) error {
	email := "fattureinvoices35@gmail.com"
	password := "FattureInvoices123456789"

	to := msg.To

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	message := []byte(msg.Subject + msg.Body)

	auth := smtp.PlainAuth("", email, password, host)
	err := smtp.SendMail(address, auth, email, to, message)
	if err != nil {
		panic(err)
	}

	return s.repo.SendEmail(msg)
}

func NewServiceSendCustomer(repo adapter.AdapterSendCustomerRepository, c config.Config) adapter.AdapterSendCustomerService {
	return &serviceSendCustomer{
		repo: repo,
		c:    c,
	}
}
