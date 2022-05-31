package usecase_test

import (
	"errors"
	"testing"

	"Back-end/config"
	"Back-end/internal/model"
	"Back-end/internal/testing/usecase_test/mocks"
	"Back-end/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsersService(t *testing.T) {
	repo := mocks.MockUserRepository{}
	userData := []model.User{
		{
			ID:       1,
			Username: "username testing 1",
			Password: "password testing 1",
		},
		{
			ID:       2,
			Username: "username testing 2",
			Password: "password testing 2",
		},
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetAllUsers").Return(userData).Once()

		svc := usecase.NewServiceUser(&repo, config.Config{})
		GetAll := svc.GetAllUsersService()

		assert.Equal(t, GetAll, userData)
	})
}

func TestGetUserByIDService(t *testing.T) {
	repo := mocks.MockUserRepository{}
	userData := model.User{
		ID:       1,
		Username: "username testing",
		Password: "password testing",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetUserByID", mock.Anything).Return(userData, nil).Once()

		svc := usecase.NewServiceUser(&repo, config.Config{})
		GetID, Err := svc.GetUserByIDService(userData.ID)

		assert.Equal(t, GetID, userData)
		assert.NoError(t, Err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetUserByID", mock.Anything).Return(userData, errors.New("Failed Get User by ID Service")).Once()

		svc := usecase.NewServiceUser(&repo, config.Config{})
		GetID, Err := svc.GetUserByIDService(userData.ID)

		assert.Equal(t, GetID, userData)
		assert.Error(t, Err)
	})
}

func TestUpdateUserByIDService(t *testing.T) {
	repo := mocks.MockUserRepository{}
	userData := model.User{
		ID:       1,
		Username: "username testing",
		Password: "password testing",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("UpdateUserByID", mock.Anything, mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceUser(&repo, config.Config{})
		Err := svc.UpdateUserByIDService(userData.ID, userData)

		assert.NoError(t, Err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("UpdateUserByID", mock.Anything, mock.Anything).Return(errors.New("Failed Update User by ID Service")).Once()

		svc := usecase.NewServiceUser(&repo, config.Config{})
		Err := svc.UpdateUserByIDService(userData.ID, userData)

		assert.Error(t, Err)
	})
}

func TestDeleteUserByIDService(t *testing.T) {
	repo := mocks.MockUserRepository{}
	userData := model.User{
		ID:       1,
		Username: "username testing",
		Password: "password testing",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("DeleteUserByID", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceUser(&repo, config.Config{})
		Err := svc.DeleteUserByIDService(userData.ID)

		assert.NoError(t, Err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("DeleteUserByID", mock.Anything).Return(errors.New("Failed Delete User by ID Service")).Once()

		svc := usecase.NewServiceUser(&repo, config.Config{})
		Err := svc.DeleteUserByIDService(userData.ID)

		assert.Error(t, Err)
	})
}
