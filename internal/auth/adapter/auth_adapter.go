package adapter

import (
	"Back-end/internal/user/model"
)

type AdapterAuthRepository interface {
	Register(user model.User) error
	UsernameExists(username string) (user model.User, err error)
	Login(username string) (user model.User, err error)
}

type AdapterAuthService interface {
	RegisterService(user model.User) (int, error)
	LoginService(username string, password string) (string, int)
}
