package usecase

import (
	"time"

	"Back-end/config"
	"Back-end/internal/user/adapter"
	"Back-end/internal/user/model"

	"golang.org/x/crypto/bcrypt"
)

type serviceUser struct {
	c    config.Config
	repo adapter.AdapterUserRepository
}

func (s *serviceUser) GetAllUsersService() []model.User {
	return s.repo.GetAllUsers()
}

func (s *serviceUser) GetUserByIDService(id int) (model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *serviceUser) UpdateUserByIDService(id int, user model.User) error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPassword)
	user.Updated_At = time.Now()

	return s.repo.UpdateUserByID(id, user)
}

func (s *serviceUser) UpdateUsernameService(id int, username string) error {
	var user model.User

	_, errUsername := s.repo.UsernameExist(username)
	if errUsername != nil {
		return errUsername
	}

	user.Username = username
	user.Updated_At = time.Now()
	return s.repo.UpdateUserByID(id, user)
}

func (s *serviceUser) UpdatePasswordService(id int, password string) error {
	var user model.User

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(hashPassword)
	user.Updated_At = time.Now()

	return s.repo.UpdateUserByID(id, user)
}

func (s *serviceUser) DeleteUserByIDService(id int) error {
	return s.repo.DeleteUserByID(id)
}

func NewServiceUser(repo adapter.AdapterUserRepository, c config.Config) adapter.AdapterUserService {
	return &serviceUser{
		repo: repo,
		c:    c,
	}
}
