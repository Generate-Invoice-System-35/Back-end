package repository

import (
	"fmt"

	"Back-end/internal/send_customer/adapter"
	"Back-end/internal/send_customer/model"

	"gorm.io/gorm"
)

type RepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *RepositoryMysqlLayer) SendEmail(message model.SendCustomer) error {
	res := r.DB.Create(&message)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) SendWhatsapp(message model.SendCustomer) error {
	message.Subject = "-"
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
