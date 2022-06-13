package repository

import (
	"fmt"

	"Back-end/internal/adapter"
	"Back-end/internal/model"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) SendEmail(message model.SendCustomer) error {
	res := r.DB.Create(&message)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func NewMysqlSendCustomerRepository(db *gorm.DB) adapter.AdapterSendCustomerRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
