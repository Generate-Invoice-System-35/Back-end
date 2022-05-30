package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

func (r *RepositoryMysqlLayer) CreateRole(role model.Role) error {
	res := r.DB.Create(&role)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllRoles() []model.Role {
	roles := []model.Role{}
	r.DB.Find(&roles)

	return roles
}

func (r *RepositoryMysqlLayer) GetRoleByID(id int) (role model.Role, err error) {
	res := r.DB.Where("id = ?", id).Find(&role)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateRoleByID(id int, role model.Role) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&role)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteRoleByID(id int) error {
	res := r.DB.Delete(&model.Role{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlRoleRepository(db *gorm.DB) adapter.AdapterRoleRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
