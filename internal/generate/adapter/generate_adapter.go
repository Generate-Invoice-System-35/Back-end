package adapter

import (
	invoice "Back-end/internal/invoice/model"
	item "Back-end/internal/invoice_item/model"
	transaction "Back-end/internal/payment_gateway/xendit/model"
	send "Back-end/internal/send_customer/model"
)

type AdapterGenerateInvoiceRepository interface {
	CreateInvoiceGenerate(invoice invoice.Invoice) error
	CreateInvoiceItemsGenerate(item item.InvoiceItem) error
	CreateTransactionRecord(int, transaction.TransactionRecord) error
	NumberInvoiceExists(number string) (invoice invoice.Invoice, flag bool)
	CreateInvoicesGenerate(invoice []invoice.Invoice, item []item.InvoiceItem) error
	GetInvoices(int) (invoice.Invoice, []item.InvoiceItem, error)
	GetTotalAmount(int) (float32, error)
	UpdateStatusInvoice(int, invoice.Invoice) error
	SendEmail(message send.SendCustomer) error
}

type AdapterGenerateInvoiceService interface {
	GenerateFileService(data [][]string) error
	GenerateInvoiceService(data []int) error
}
