package usecase

import (
	"net/http"

	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/helper"
	"Back-end/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type serviceAuth struct {
	c    config.Config
	repo adapter.AdapterAuthRepository
}

func (s *serviceAuth) RegisterService(user model.User) (error, int) {
	_, errUsername := s.repo.UsernameExists(user.Username)
	if errUsername != nil {
		return errUsername, http.StatusExpectationFailed
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPassword)

	errRegister := s.repo.Register(user)
	if errRegister != nil {
		return errRegister, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *serviceAuth) LoginService(username string, password string) (string, int) {
	user, err := s.repo.Login(username)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errPass != nil {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(user.ID, user.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func NewServiceAuth(repo adapter.AdapterAuthRepository, c config.Config) adapter.AdapterAuthService {
	return &serviceAuth{
		repo: repo,
		c:    c,
	}
}
