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
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterService(t *testing.T) {
	repo := mocks.MockAuthRepository{}
	userData := model.User{
		ID:       1,
		Username: "username testing",
		Password: "password testing",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(nil).Once()
		repo.On("UsernameExists", mock.Anything).Return(userData, nil).Once()

		svc := usecase.NewServiceAuth(&repo, config.Config{})
		Err, Status := svc.RegisterService(userData)

		assert.Equal(t, 200, Status)
		assert.NoError(t, Err)
	})

	t.Run("Failed Expectation Failed", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(errors.New("Failed Register")).Once()
		repo.On("UsernameExists", mock.Anything).Return(userData, errors.New("Failed Username Exist")).Once()

		svc := usecase.NewServiceAuth(&repo, config.Config{})
		Err, Status := svc.RegisterService(userData)

		assert.Equal(t, 417, Status)
		assert.Error(t, Err)
	})
}

func TestLoginService(t *testing.T) {
	repo := mocks.MockAuthRepository{}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("password testing"), 14)
	userData := model.User{
		ID:       1,
		Username: "username testing",
		Password: string(hashPassword),
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(userData, nil).Once()

		svc := usecase.NewServiceAuth(&repo, config.Config{})
		_, status := svc.LoginService(userData.Username, "password testing")

		assert.Equal(t, 200, status)
	})

	t.Run("Failed Unauthorized", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(userData, nil).Once()

		svc := usecase.NewServiceAuth(&repo, config.Config{})
		_, status := svc.LoginService(userData.Username, "Password Gagal")

		assert.Equal(t, 401, status)
	})

	t.Run("Failed Internal Server", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(userData, errors.New("Failed Internal Server Login")).Once()

		svc := usecase.NewServiceAuth(&repo, config.Config{})
		_, status := svc.LoginService(userData.Username, userData.Password)

		assert.Equal(t, 500, status)
	})
}
