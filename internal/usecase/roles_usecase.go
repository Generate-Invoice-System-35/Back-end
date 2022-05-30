package usecase

import (
	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

type serviceRole struct {
	c    config.Config
	repo adapter.AdapterRoleRepository
}

func (s *serviceRole) CreateRoleService(role model.Role) error {
	return s.repo.CreateRole(role)
}

func (s *serviceRole) GetAllRolesService() []model.Role {
	return s.repo.GetAllRoles()
}

func (s *serviceRole) GetRoleByIDService(id int) (model.Role, error) {
	return s.repo.GetRoleByID(id)
}

func (s *serviceRole) UpdateRoleByIDService(id int, role model.Role) error {
	return s.repo.UpdateRoleByID(id, role)
}

func (s *serviceRole) DeleteRoleByIDService(id int) error {
	return s.repo.DeleteRoleByID(id)
}

func NewServiceRole(repo adapter.AdapterRoleRepository, c config.Config) adapter.AdapterRoleService {
	return &serviceRole{
		repo: repo,
		c:    c,
	}
}
