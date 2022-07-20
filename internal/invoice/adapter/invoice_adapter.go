package adapter

import "Back-end/internal/invoice/model"

type AdapterInvoiceRepository interface {
	CreateInvoice(invoice model.Invoice) error
	GetAllInvoices() []model.Invoice
	GetInvoiceByID(id int) (invoice model.Invoice, err error)
	GetInvoicesPagination(pagination model.Pagination) ([]model.Invoice, error)
	GetTotalPagesPagination() (int, error)
	GetInvoicesByPaymentStatus(status int, pagination model.Pagination) (invoice []model.Invoice)
	GetInvoicesByNameCustomer(name string) (invoice []model.Invoice, err error)
	UpdateInvoiceByID(id int, invoice model.Invoice) error
	DeleteInvoiceByID(id int) error
}

type AdapterInvoiceService interface {
	CreateInvoiceService(invoice model.Invoice) error
	GetAllInvoicesService() []model.Invoice
	GetInvoiceByIDService(id int) (invoice model.Invoice, err error)
	GetInvoicesPaginationService(pagination model.Pagination) ([]model.Invoice, error)
	GetTotalPagesPaginationService() (int, error)
	GetInovicesByPaymentStatusService(status int, pagination model.Pagination) (invoice []model.Invoice)
	GetInvoicesByNameCustomerService(name string) (invoice []model.Invoice, err error)
	UpdateInvoiceByIDService(id int, invoice model.Invoice) error
	DeleteInvoiceByIDService(id int) error
}
