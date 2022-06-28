package adapter

import (
	"Back-end/internal/model"

	"github.com/xendit/xendit-go"
)

type AdapterPaymentGatewayRepository interface {
	GetInvoices(int) (model.Invoice, model.InvoiceItem, error)
}

type AdapterPaymentGatewayService interface {
	CreateInvoiceService(int) (*xendit.Invoice, error)
}
