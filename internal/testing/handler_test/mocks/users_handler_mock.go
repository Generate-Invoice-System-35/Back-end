package mocks

import (
	"Back-end/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (r *MockUserService) GetAllUsersService() []model.User {
	ret := r.Called()

	var user []model.User
	if res, ok := ret.Get(0).(func() []model.User); ok {
		user = res()
	} else {
		if ret.Get(0) != nil {
			user = ret.Get(0).([]model.User)
		}
	}

	return user
}

func (r *MockUserService) GetUserByIDService(id int) (model.User, error) {
	ret := r.Called(id)

	var user model.User
	if res, ok := ret.Get(0).(func(int) model.User); ok {
		user = res(id)
	} else {
		user = ret.Get(0).(model.User)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return user, err
}

func (r *MockUserService) UpdateUserByIDService(id int, user model.User) error {
	ret := r.Called(id, user)

	var err error
	if res, ok := ret.Get(0).(func(int, model.User) error); ok {
		err = res(id, user)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockUserService) DeleteUserByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
