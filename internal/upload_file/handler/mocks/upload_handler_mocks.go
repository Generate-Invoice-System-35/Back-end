package mocks

import (
	"mime/multipart"

	"Back-end/internal/upload_file/model"

	"github.com/stretchr/testify/mock"
)

type MockUploadService struct {
	mock.Mock
}

func (r *MockUploadService) CreateImageService(image model.File, file *multipart.FileHeader) error {
	ret := r.Called(image, file)

	var err error
	if res, ok := ret.Get(0).(func(model.File, *multipart.FileHeader) error); ok {
		err = res(image, file)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockUploadService) GetAllImagesService() []model.File {
	ret := r.Called()

	var files []model.File
	if res, ok := ret.Get(0).(func() []model.File); ok {
		files = res()
	} else {
		if ret.Get(0) != nil {
			files = ret.Get(0).([]model.File)
		}
	}

	return files
}

func (r *MockUploadService) GetImageByIDService(id int) (image model.File, err error) {
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

func (r *MockUploadService) UpdateImageService(id int, image model.File, file *multipart.FileHeader) error {
	ret := r.Called(id, image, file)

	var err error
	if res, ok := ret.Get(0).(func(int, model.File, *multipart.FileHeader) error); ok {
		err = res(id, image, file)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockUploadService) DeleteImageService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
