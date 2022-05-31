package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

func (r *RepositoryMysqlLayer) Register(user model.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) Login(username string) (user model.User, err error) {
	res := r.DB.Where("username = ?", username).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("username not found")
	}

	return
}

func NewMysqlAuthRepository(db *gorm.DB) adapter.AdapterAuthRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
