package mocks

import (
	"Back-end/internal/send_customer/model"

	"github.com/stretchr/testify/mock"
)

type MockSendCustomerRepository struct {
	mock.Mock
}

func (r *MockSendCustomerRepository) SendEmail(message model.SendCustomer) error {
	ret := r.Called(message)

	var err error
	if res, ok := ret.Get(0).(func(model.SendCustomer) error); ok {
		err = res(message)
	} else {
		err = ret.Error(0)
	}

	return err
}
