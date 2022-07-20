package mocks

import (
	"Back-end/internal/invoice/model"

	"github.com/stretchr/testify/mock"
)

type MockInvoiceService struct {
	mock.Mock
}

func (r *MockInvoiceService) CreateInvoiceService(invoice model.Invoice) error {
	ret := r.Called(invoice)

	var err error
	if res, ok := ret.Get(0).(func(model.Invoice) error); ok {
		err = res(invoice)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockInvoiceService) GetAllInvoicesService() []model.Invoice {
	ret := r.Called()

	var invoices []model.Invoice
	if res, ok := ret.Get(0).(func() []model.Invoice); ok {
		invoices = res()
	} else {
		invoices = ret.Get(0).([]model.Invoice)
	}

	return invoices
}

func (r *MockInvoiceService) GetInvoiceByIDService(id int) (invoice model.Invoice, err error) {
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

func (r *MockInvoiceService) GetInvoicesPaginationService(pagination model.Pagination) ([]model.Invoice, error) {
	ret := r.Called(pagination)

	var invoices []model.Invoice
	if res, ok := ret.Get(0).(func(model.Pagination) []model.Invoice); ok {
		invoices = res(pagination)
	} else {
		invoices = ret.Get(0).([]model.Invoice)
	}

	var err error
	if res, ok := ret.Get(1).(func(model.Pagination) error); ok {
		err = res(pagination)
	} else {
		err = ret.Error(1)
	}

	return invoices, err
}

func (r *MockInvoiceService) GetInovicesByPaymentStatusService(status int) (invoice []model.Invoice, err error) {
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

func (r *MockInvoiceService) GetInvoicesByNameCustomerService(name string) (invoice []model.Invoice, err error) {
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

func (r *MockInvoiceService) UpdateInvoiceByIDService(id int, invoice model.Invoice) error {
	ret := r.Called(id, invoice)

	var err error
	if res, ok := ret.Get(0).(func(int, model.Invoice) error); ok {
		err = res(id, invoice)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockInvoiceService) DeleteInvoiceByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
