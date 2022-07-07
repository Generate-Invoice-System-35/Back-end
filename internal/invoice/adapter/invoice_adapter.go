package adapter

import "Back-end/internal/invoice/model"

type AdapterInvoiceRepository interface {
	CreateInvoice(invoice model.Invoice) error
	GetAllInvoices() []model.Invoice
	GetInvoiceByID(id int) (invoice model.Invoice, err error)
	GetInvoicesByPaymentStatus(status int) (invoice []model.Invoice, err error)
	GetInvoicesByNameCustomer(name string) (invoice []model.Invoice, err error)
	UpdateInvoiceByID(id int, invoice model.Invoice) error
	DeleteInvoiceByID(id int) error
}

type AdapterInvoiceService interface {
	CreateInvoiceService(invoice model.Invoice) error
	GetAllInvoicesService() []model.Invoice
	GetInvoiceByIDService(id int) (invoice model.Invoice, err error)
	GetInovicesByPaymentStatusService(status int) (invoice []model.Invoice, err error)
	GetInvoicesByNameCustomerService(name string) (invoice []model.Invoice, err error)
	UpdateInvoiceByIDService(id int, invoice model.Invoice) error
	DeleteInvoiceByIDService(id int) error
}
