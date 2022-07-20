package mocks

import "github.com/stretchr/testify/mock"

type MockGenerateService struct {
	mock.Mock
}

func (r *MockGenerateService) GenerateFileService(data [][]string) error {
	ret := r.Called(data)

	var err error
	if res, ok := ret.Get(0).(func([][]string) error); ok {
		err = res(data)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockGenerateService) GenerateInvoiceService(data []int) error {
	ret := r.Called(data)

	var err error
	if res, ok := ret.Get(0).(func([]int) error); ok {
		err = res(data)
	} else {
		err = ret.Error(0)
	}

	return err
}
