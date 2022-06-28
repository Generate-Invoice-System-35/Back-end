package mysql

import (
	"Back-end/internal/model"
	"Back-end/internal/payment_gateway/adapter"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type RepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *RepositoryMysqlLayer) GetInvoices(id int) (inv model.Invoice, items model.InvoiceItem, err error) {
	res1 := r.DB.Where("id = ?", id).Find(&inv)
	if res1.RowsAffected < 1 {
		log.Printf("not found invoice")
		err = fmt.Errorf("not found invoice")
	}

	res2 := r.DB.Where("id_invoice = ?", id).Find(&items)
	if res2.RowsAffected < 1 {
		log.Printf("not found invoice items")
		err = fmt.Errorf("not found invoice items")
	}

	return
}

func NewMysqlPaymentGatewayRepository(db *gorm.DB) adapter.AdapterPaymentGatewayRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
