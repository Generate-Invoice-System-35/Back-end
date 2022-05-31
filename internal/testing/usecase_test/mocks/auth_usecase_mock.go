package mocks

import (
	"Back-end/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	mock.Mock
}

func (r *MockAuthRepository) Register(user model.User) error {
	ret := r.Called(user)

	var err error
	if res, ok := ret.Get(0).(func(model.User) error); ok {
		err = res(user)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockAuthRepository) Login(username string) (model.User, error) {
	ret := r.Called(username)

	var user model.User
	if res, ok := ret.Get(0).(func(string) model.User); ok {
		user = res(username)
	} else {
		user = ret.Get(0).(model.User)
	}

	var err error
	if res, ok := ret.Get(1).(func(string) error); ok {
		err = res(username)
	} else {
		err = ret.Error(1)
	}

	return user, err
}
