package adapter

import (
	invoice "Back-end/internal/invoice/model"
	item "Back-end/internal/invoice_item/model"
	transaction "Back-end/internal/payment_gateway/xendit/model"

	"github.com/xendit/xendit-go"
)

type AdapterPaymentGatewayRepository interface {
	CreateTransactionRecord(int, transaction.TransactionRecord) error
	GetIDInvoicePayment(int) (record transaction.TransactionRecord, err error)
	GetInvoices(int) (invoice.Invoice, []item.InvoiceItem, error)
	GetTotalAmount(int) (float32, error)
	UpdateStatusInvoice(int, invoice.Invoice) error
}

type AdapterPaymentGatewayService interface {
	CreateXenditPaymentInvoiceService(int) (*xendit.Invoice, error)
	GetXenditPaymentInvoiceService(int) (*xendit.Invoice, error)
	GetAllXenditPaymentInvoiceService() ([]xendit.Invoice, error)
	ExpireXenditPaymentInvoiceService(int) (*xendit.Invoice, error)
	CallbackXenditPaymentInvoiceService(transaction.CallbackInvoice) error
}
