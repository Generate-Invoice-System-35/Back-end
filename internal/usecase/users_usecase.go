package usecase

import (
	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"

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
