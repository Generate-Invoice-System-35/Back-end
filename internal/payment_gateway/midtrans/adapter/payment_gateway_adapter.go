package adapter

import "Back-end/internal/model"

type AdapterPaymentGatewayRepository interface {
	GetInvoices(int) (model.Invoice, []model.InvoiceItem, error)
	GetTotalAmount(int) (float32, error)
}

type AdapterPaymentGatewayService interface {
	ChargeTransactionService(int) (int, error)
}
