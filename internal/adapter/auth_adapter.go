package adapter

import (
	"Back-end/internal/model"
)

type AdapterAuthRepository interface {
	Register(user model.User) error
	Login(username string) (user model.User, err error)
}

type AdapterAuthService interface {
	RegisterService(user model.User) error
	LoginService(username string, password string) (string, int)
}
