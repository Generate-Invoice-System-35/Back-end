package mocks

import (
	"Back-end/internal/invoice_payment_status/model"

	"github.com/stretchr/testify/mock"
)

type MockInvoicePaymentStatusService struct {
	mock.Mock
}

func (r *MockInvoicePaymentStatusService) CreateInvoicePaymentStatusService(IPStatus model.InvoicePaymentStatus) error {
	ret := r.Called(IPStatus)

	var err error
	if res, ok := ret.Get(0).(func(model.InvoicePaymentStatus) error); ok {
		err = res(IPStatus)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockInvoicePaymentStatusService) GetAllInvoicesPaymentStatusService() []model.InvoicePaymentStatus {
	ret := r.Called()

	var status []model.InvoicePaymentStatus
	if res, ok := ret.Get(0).(func() []model.InvoicePaymentStatus); ok {
		status = res()
	} else {
		status = ret.Get(0).([]model.InvoicePaymentStatus)
	}

	return status
}

func (r *MockInvoicePaymentStatusService) GetInvoicePaymentStatusByIDService(id int) (IPStatus model.InvoicePaymentStatus, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) model.InvoicePaymentStatus); ok {
		IPStatus = res(id)
	} else {
		IPStatus = ret.Get(0).(model.InvoicePaymentStatus)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockInvoicePaymentStatusService) UpdateInvoicePaymentStatusByIDService(id int, IPStatus model.InvoicePaymentStatus) error {
	ret := r.Called(id, IPStatus)

	var err error
	if res, ok := ret.Get(0).(func(int, model.InvoicePaymentStatus) error); ok {
		err = res(id, IPStatus)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockInvoicePaymentStatusService) DeleteInvoicePaymentStatusByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
