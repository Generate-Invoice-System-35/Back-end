package adapter

import "Back-end/internal/model"

type AdapterPaymentStatusRepository interface {
	CreateInvoicePaymentStatus(IPStatus model.InvoicePaymentStatus) error
	GetAllInvoicesPaymentStatus() []model.InvoicePaymentStatus
	GetInvoicePaymentStatusByID(id int) (IPStatus model.InvoicePaymentStatus, err error)
	UpdateInvoicePaymentStatusByID(id int, IPStatus model.InvoicePaymentStatus) error
	DeleteInvoicePaymentStatusByID(id int) error
}

type AdapterPaymentStatusService interface {
	CreateInvoicePaymentStatusService(IPStatus model.InvoicePaymentStatus) error
	GetAllInvoicesPaymentStatusService() []model.InvoicePaymentStatus
	GetInvoicePaymentStatusByIDService(id int) (IPStatus model.InvoicePaymentStatus, err error)
	UpdateInvoicePaymentStatusByIDService(id int, IPStatus model.InvoicePaymentStatus) error
	DeleteInvoicePaymentStatusByIDService(id int) error
}
