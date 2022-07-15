package mocks

import (
	"Back-end/internal/upload_file/model"

	"github.com/stretchr/testify/mock"
)

type MockImageRepository struct {
	mock.Mock
}

func (r *MockImageRepository) CreateImage(image model.File) error {
	ret := r.Called(image)

	var err error
	if res, ok := ret.Get(0).(func(model.File) error); ok {
		err = res(image)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockImageRepository) GetAllImages() []model.File {
	ret := r.Called()

	var images []model.File
	if res, ok := ret.Get(0).(func() []model.File); ok {
		images = res()
	} else {
		if ret.Get(0) != nil {
			images = ret.Get(0).([]model.File)
		}
	}

	return images
}

func (r *MockImageRepository) GetImageByID(id int) (image model.File, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) model.File); ok {
		image = res(id)
	} else {
		image = ret.Get(0).(model.File)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockImageRepository) UpdateImage(id int, image model.File) error {
	ret := r.Called(id, image)

	var err error
	if res, ok := ret.Get(0).(func(int, model.File) error); ok {
		err = res(id, image)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockImageRepository) DeleteImage(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
