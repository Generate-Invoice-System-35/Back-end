package mocks

import (
	"Back-end/internal/invoice/model"

	"github.com/stretchr/testify/mock"
)

type MockInvoiceRepository struct {
	mock.Mock
}

func (r *MockInvoiceRepository) CreateInvoice(invoice model.Invoice) error {
	ret := r.Called(invoice)

	var err error
	if res, ok := ret.Get(0).(func(model.Invoice) error); ok {
		err = res(invoice)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockInvoiceRepository) GetAllInvoices() []model.Invoice {
	ret := r.Called()

	var invoices []model.Invoice
	if res, ok := ret.Get(0).(func() []model.Invoice); ok {
		invoices = res()
	} else {
		if ret.Get(0) != nil {
			invoices = ret.Get(0).([]model.Invoice)
		}
	}

	return invoices
}

func (r *MockInvoiceRepository) GetInvoiceByID(id int) (invoice model.Invoice, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) model.Invoice); ok {
		invoice = res(id)
	} else {
		invoice = ret.Get(0).(model.Invoice)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockInvoiceRepository) GetInvoicesPagination(pagination model.Pagination) (invoice []model.Invoice, err error) {
	ret := r.Called(pagination)

	if res, ok := ret.Get(0).(func(model.Pagination) []model.Invoice); ok {
		invoice = res(pagination)
	} else {
		invoice = ret.Get(0).([]model.Invoice)
	}

	if res, ok := ret.Get(1).(func(model.Pagination) error); ok {
		err = res(pagination)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockInvoiceRepository) GetInvoicesByPaymentStatus(status int) (invoice []model.Invoice, err error) {
	ret := r.Called(status)

	if res, ok := ret.Get(0).(func(int) []model.Invoice); ok {
		invoice = res(status)
	} else {
		invoice = ret.Get(0).([]model.Invoice)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(status)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockInvoiceRepository) GetInvoicesByNameCustomer(name string) (invoice []model.Invoice, err error) {
	ret := r.Called(name)

	if res, ok := ret.Get(0).(func(string) []model.Invoice); ok {
		invoice = res(name)
	} else {
		invoice = ret.Get(0).([]model.Invoice)
	}

	if res, ok := ret.Get(1).(func(string) error); ok {
		err = res(name)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockInvoiceRepository) UpdateInvoiceByID(id int, invoice model.Invoice) error {
	ret := r.Called(id, invoice)

	var err error
	if res, ok := ret.Get(0).(func(int, model.Invoice) error); ok {
		err = res(id, invoice)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockInvoiceRepository) DeleteInvoiceByID(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
