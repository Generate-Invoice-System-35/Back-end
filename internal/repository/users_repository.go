package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

type RepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *RepositoryMysqlLayer) GetAllUsers() []model.User {
	users := []model.User{}
	r.DB.Find(&users)

	return users
}

func (r *RepositoryMysqlLayer) GetUserByID(id int) (user model.User, err error) {
	res := r.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UsernameExist(username string) (user model.User, err error) {
	res := r.DB.Where("username = ?", username).Find(&user)
	if res.RowsAffected > 0 {
		err = fmt.Errorf("username exist")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateUserByID(id int, user model.User) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteUserByID(id int) error {
	res := r.DB.Delete(&model.User{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlUserRepository(db *gorm.DB) adapter.AdapterUserRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
