package mocks

import (
	"Back-end/internal/invoice_item/model"

	"github.com/stretchr/testify/mock"
)

type MockInvoiceItemService struct {
	mock.Mock
}

func (r *MockInvoiceItemService) CreateInvoiceItemService(item model.InvoiceItem) error {
	ret := r.Called(item)

	var err error
	if res, ok := ret.Get(0).(func(model.InvoiceItem) error); ok {
		err = res(item)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockInvoiceItemService) GetAllInvoiceItemsService() []model.InvoiceItem {
	ret := r.Called()

	var items []model.InvoiceItem
	if res, ok := ret.Get(0).(func() []model.InvoiceItem); ok {
		items = res()
	} else {
		items = ret.Get(0).([]model.InvoiceItem)
	}

	return items
}

func (r *MockInvoiceItemService) GetInvoiceItemByIDService(id int) (invItem model.InvoiceItem, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) model.InvoiceItem); ok {
		invItem = res(id)
	} else {
		invItem = ret.Get(0).(model.InvoiceItem)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockInvoiceItemService) GetInvoiceItemByNumberService(number string) (invItem []model.InvoiceItem, err error) {
	ret := r.Called(number)

	if res, ok := ret.Get(0).(func(string) []model.InvoiceItem); ok {
		invItem = res(number)
	} else {
		invItem = ret.Get(0).([]model.InvoiceItem)
	}

	if res, ok := ret.Get(1).(func(string) error); ok {
		err = res(number)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockInvoiceItemService) UpdateInvoiceItemByIDService(id int, invItem model.InvoiceItem) error {
	ret := r.Called(id, invItem)

	var err error
	if res, ok := ret.Get(0).(func(int, model.InvoiceItem) error); ok {
		err = res(id, invItem)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockInvoiceItemService) DeleteInvoiceItemByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
