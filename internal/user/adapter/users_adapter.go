package adapter

import "Back-end/internal/user/model"

type AdapterUserRepository interface {
	GetAllUsers() []model.User
	GetUserByID(id int) (user model.User, err error)
	UsernameExist(username string) (user model.User, err error)
	UpdateUserByID(id int, user model.User) error
	DeleteUserByID(id int) error
}

type AdapterUserService interface {
	GetAllUsersService() []model.User
	GetUserByIDService(id int) (model.User, error)
	UpdateUserByIDService(id int, user model.User) error
	UpdateUsernameService(id int, username string) error
	UpdatePasswordService(id int, password string) error
	DeleteUserByIDService(id int) error
}
