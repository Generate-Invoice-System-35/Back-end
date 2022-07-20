package mocks

import (
	"Back-end/internal/user/model"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (r *MockUserRepository) GetAllUsers() []model.User {
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

func (r *MockUserRepository) GetUserByID(id int) (user model.User, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) model.User); ok {
		user = res(id)
	} else {
		user = ret.Get(0).(model.User)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockUserRepository) UsernameExist(username string) (user model.User, err error) {
	ret := r.Called(username)

	if res, ok := ret.Get(0).(func(string) model.User); ok {
		user = res(username)
	} else {
		user = ret.Get(0).(model.User)
	}

	if res, ok := ret.Get(1).(func(string) error); ok {
		err = res(username)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockUserRepository) UpdateUserByID(id int, user model.User) error {
	ret := r.Called(id, user)

	var err error
	if res, ok := ret.Get(0).(func(int, model.User) error); ok {
		err = res(id, user)
	} else {
		err = ret.Error(0)
	}

	return err
}

// func (r *MockUserRepository) UpdateUsernameService(id int, username string) error {
// 	ret := r.Called(id)

// 	var err error
// 	if res, ok := ret.Get(0).(func(int, string) error); ok {
// 		err = res(id, username)
// 	} else {
// 		err = ret.Error(0)
// 	}

// 	return err
// }

// func (r *MockUserRepository) UpdatePasswordService(id int, password string) error {
// 	ret := r.Called(id)

// 	var err error
// 	if res, ok := ret.Get(0).(func(int, string) error); ok {
// 		err = res(id, password)
// 	} else {
// 		err = ret.Error(0)
// 	}

// 	return err
// }

func (r *MockUserRepository) DeleteUserByID(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
