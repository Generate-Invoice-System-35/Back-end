package mocks

import (
	invoices "Back-end/internal/invoice/model"
	items "Back-end/internal/invoice_item/model"
	transactions "Back-end/internal/payment_gateway/xendit/model"
	sends "Back-end/internal/send_customer/model"

	"github.com/stretchr/testify/mock"
)

type MockGenerateRepository struct {
	mock.Mock
}

func (r *MockGenerateRepository) CreateInvoiceGenerate(invoice invoices.Invoice) error {
	ret := r.Called(invoice)

	var err error
	if res, ok := ret.Get(0).(func(invoices.Invoice) error); ok {
		err = res(invoice)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockGenerateRepository) CreateInvoiceItemsGenerate(item items.InvoiceItem) error {
	ret := r.Called(item)

	var err error
	if res, ok := ret.Get(0).(func(items.InvoiceItem) error); ok {
		err = res(item)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockGenerateRepository) CreateTransactionRecord(id int, transaction transactions.TransactionRecord) error {
	ret := r.Called(transaction)

	var err error
	if res, ok := ret.Get(0).(func(transactions.TransactionRecord) error); ok {
		err = res(transaction)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockGenerateRepository) NumberInvoiceExists(number string) (invoice invoices.Invoice, flag bool) {
	ret := r.Called(number)

	if res, ok := ret.Get(0).(func(string) invoices.Invoice); ok {
		invoice = res(number)
	} else {
		invoice = ret.Get(0).(invoices.Invoice)
	}

	if res, ok := ret.Get(1).(func(string) bool); ok {
		flag = res(number)
	} else {
		flag = ret.Get(1).(bool)
	}

	return
}

func (r *MockGenerateRepository) CreateInvoicesGenerate(invoice []invoices.Invoice, item []items.InvoiceItem) error {
	ret := r.Called(invoice, item)

	var err error
	if res, ok := ret.Get(0).(func([]invoices.Invoice, []items.InvoiceItem) error); ok {
		err = res(invoice, item)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockGenerateRepository) GetInvoices(id int) (invoice invoices.Invoice, item []items.InvoiceItem, err error) {
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

func (r *MockGenerateRepository) GetTotalAmount(id int) (total float32, err error) {
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

func (r *MockGenerateRepository) UpdateStatusInvoice(id int, invoice invoices.Invoice) error {
	ret := r.Called(id, invoice)

	var err error
	if res, ok := ret.Get(0).(func(int, invoices.Invoice) error); ok {
		err = res(id, invoice)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockGenerateRepository) SendEmail(message sends.SendCustomer) error {
	ret := r.Called(message)

	var err error
	if res, ok := ret.Get(0).(func(sends.SendCustomer) error); ok {
		err = res(message)
	} else {
		err = ret.Error(0)
	}

	return err
}
