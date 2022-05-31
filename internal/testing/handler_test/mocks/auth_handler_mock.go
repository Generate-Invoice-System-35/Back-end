package mocks

import (
	"Back-end/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (r *MockAuthService) RegisterService(user model.User) error {
	ret := r.Called(user)

	var err error
	if res, ok := ret.Get(0).(func(model.User) error); ok {
		err = res(user)
	} else {
		err = ret.Error(0)
	}

	return err
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
