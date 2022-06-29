package adapter

import (
	"Back-end/internal/model"

	"github.com/xendit/xendit-go"
)

type AdapterPaymentGatewayRepository interface {
	CreateTransactionRecord(int, model.TransactionRecord) error
	GetIDInvoicePayment(int) (record model.TransactionRecord, err error)
	GetInvoices(int) (model.Invoice, []model.InvoiceItem, error)
	GetTotalAmount(int) (float32, error)
	UpdateStatusInvoice(int, model.Invoice) error
}

type AdapterPaymentGatewayService interface {
	CreateXenditPaymentInvoiceService(int) (*xendit.Invoice, error)
	GetXenditPaymentInvoiceService(int) (*xendit.Invoice, error)
	GetAllXenditPaymentInvoiceService() ([]xendit.Invoice, error)
	ExpireXenditPaymentInvoiceService(int) (*xendit.Invoice, error)
	CallbackXenditPaymentInvoiceService(model.CallbackInvoice) error
}
