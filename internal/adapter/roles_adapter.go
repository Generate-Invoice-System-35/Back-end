package adapter

import "Back-end/internal/model"

type AdapterRoleRepository interface {
	CreateRole(role model.Role) error
	GetAllRoles() []model.Role
	GetRoleByID(id int) (role model.Role, err error)
	UpdateRoleByID(id int, role model.Role) error
	DeleteRoleByID(id int) error
}

type AdapterRoleService interface {
	CreateRoleService(role model.Role) error
	GetAllRolesService() []model.Role
	GetRoleByIDService(id int) (model.Role, error)
	UpdateRoleByIDService(id int, role model.Role) error
	DeleteRoleByIDService(id int) error
}
