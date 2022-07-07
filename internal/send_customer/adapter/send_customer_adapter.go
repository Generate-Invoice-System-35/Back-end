package adapter

import "Back-end/internal/send_customer/model"

type AdapterSendCustomerRepository interface {
	SendEmail(message model.SendCustomer) error
	SendWhatsapp(message model.SendCustomer) error
}

type AdapterSendCustomerService interface {
	SendEmailService(message model.SendCustomer) error
	SendWhatsappService(message model.SendCustomer) error
}
