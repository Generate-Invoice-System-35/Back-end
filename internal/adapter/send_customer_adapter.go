package adapter

import "Back-end/internal/model"

type AdapterSendCustomerRepository interface {
	SendEmail(message model.SendCustomer) error
}

type AdapterSendCustomerService interface {
	SendEmailService(message model.SendCustomer) error
}
