package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

func (r *RepositoryMysqlLayer) UploadImage(image model.File) error {
	res := r.DB.Create(&image)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func NewMysqlUploadRepository(db *gorm.DB) adapter.AdapterUploadRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
