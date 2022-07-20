package mocks

import (
	"Back-end/internal/payment_gateway/xendit/model"

	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
)

type MockPaymentGatewayXenditService struct {
	mock.Mock
}

func (r *MockPaymentGatewayXenditService) CreateXenditPaymentInvoiceService(id int) (*xendit.Invoice, error) {
	ret := r.Called(id)

	var invoice *xendit.Invoice
	if res, ok := ret.Get(0).(func(int) *xendit.Invoice); ok {
		invoice = res(id)
	} else {
		invoice = ret.Get(0).(*xendit.Invoice)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return invoice, err
}

func (r *MockPaymentGatewayXenditService) GetXenditPaymentInvoiceService(id int) (*xendit.Invoice, error) {
	ret := r.Called(id)

	var invoice *xendit.Invoice
	if res, ok := ret.Get(0).(func(int) *xendit.Invoice); ok {
		invoice = res(id)
	} else {
		invoice = ret.Get(0).(*xendit.Invoice)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return invoice, err
}

func (r *MockPaymentGatewayXenditService) GetAllXenditPaymentInvoiceService() ([]xendit.Invoice, error) {
	ret := r.Called()

	var invoices []xendit.Invoice
	if res, ok := ret.Get(0).(func() []xendit.Invoice); ok {
		invoices = res()
	} else {
		invoices = ret.Get(0).([]xendit.Invoice)
	}

	var err error
	if res, ok := ret.Get(1).(func() error); ok {
		err = res()
	} else {
		err = ret.Error(1)
	}

	return invoices, err
}

func (r *MockPaymentGatewayXenditService) ExpireXenditPaymentInvoiceService(id int) (*xendit.Invoice, error) {
	ret := r.Called(id)

	var invoice *xendit.Invoice
	if res, ok := ret.Get(0).(func(int) *xendit.Invoice); ok {
		invoice = res(id)
	} else {
		invoice = ret.Get(0).(*xendit.Invoice)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return invoice, err
}

func (r *MockPaymentGatewayXenditService) CallbackXenditPaymentInvoiceService(callback model.CallbackInvoice) error {
	ret := r.Called(callback)

	var err error
	if res, ok := ret.Get(0).(func(model.CallbackInvoice) error); ok {
		err = res(callback)
	} else {
		err = ret.Error(0)
	}

	return err
}
