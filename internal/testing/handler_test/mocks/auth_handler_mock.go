package mocks

import (
	"Back-end/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (r *MockAuthService) RegisterService(user model.User) (int, error) {
	ret := r.Called(user)

	var status int
	if res, ok := ret.Get(0).(func(model.User) int); ok {
		status = res(user)
	} else {
		status = ret.Get(0).(int)
	}

	var err error
	if res, ok := ret.Get(1).(func(model.User) error); ok {
		err = res(user)
	} else {
		err = ret.Error(1)
	}

	return status, err
}

func (r *MockAuthService) LoginService(username string, password string) (string, int) {
	ret := r.Called(username, password)

	var token string
	if res, ok := ret.Get(0).(func(string, string) string); ok {
		token = res(username, password)
	} else {
		token = ret.Get(0).(string)
	}

	var status int
	if res, ok := ret.Get(1).(func(string, string) int); ok {
		status = res(username, password)
	} else {
		status = ret.Get(1).(int)
	}

	return token, status
}
