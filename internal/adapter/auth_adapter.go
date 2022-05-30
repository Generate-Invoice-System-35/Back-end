package adapter

import "Back-end/internal/model"

type AdapterAuthRepository interface {
	Register(user model.User, card model.Card) error
	Login(email string, password string) (string, int)
}

type AdapterAuthService interface {
	RegisterService(user model.User, card model.Card) error
	LoginService(email string, password string) (string, int)
}
