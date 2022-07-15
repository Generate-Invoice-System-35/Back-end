package mocks

import (
	invoices "Back-end/internal/invoice/model"
	items "Back-end/internal/invoice_item/model"
	transactions "Back-end/internal/payment_gateway/xendit/model"

	"github.com/stretchr/testify/mock"
)

type MockPaymentGatewayRepository struct {
	mock.Mock
}

func (r *MockPaymentGatewayRepository) CreateTransactionRecord(id int, transaction transactions.TransactionRecord) error {
	ret := r.Called(id, transaction)

	var err error
	if res, ok := ret.Get(0).(func(int, transactions.TransactionRecord) error); ok {
		err = res(id, transaction)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockPaymentGatewayRepository) GetIDInvoicePayment(id int) (record transactions.TransactionRecord, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) transactions.TransactionRecord); ok {
		record = res(id)
	} else {
		record = ret.Get(0).(transactions.TransactionRecord)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockPaymentGatewayRepository) GetInvoices(id int) (invoice invoices.Invoice, item []items.InvoiceItem, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) invoices.Invoice); ok {
		invoice = res(id)
	} else {
		invoice = ret.Get(0).(invoices.Invoice)
	}

	if res, ok := ret.Get(1).(func(int) []items.InvoiceItem); ok {
		item = res(id)
	} else {
		item = ret.Get(1).([]items.InvoiceItem)
	}

	if res, ok := ret.Get(2).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(2)
	}

	return
}

func (r *MockPaymentGatewayRepository) GetTotalAmount(id int) (total float32, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) float32); ok {
		total = res(id)
	} else {
		total = ret.Get(0).(float32)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockPaymentGatewayRepository) UpdateStatusInvoice(id int, invoice invoices.Invoice) error {
	ret := r.Called(id, invoice)

	var err error
	if res, ok := ret.Get(0).(func(int, invoices.Invoice) error); ok {
		err = res(id, invoice)
	} else {
		err = ret.Error(0)
	}

	return err
}
