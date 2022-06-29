package mysql

import (
	"fmt"
	"log"

	"Back-end/internal/model"
	"Back-end/internal/payment_gateway/midtrans/adapter"

	"gorm.io/gorm"
)

type RepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *RepositoryMysqlLayer) GetInvoices(id int) (inv model.Invoice, items []model.InvoiceItem, err error) {
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

func (r *RepositoryMysqlLayer) GetTotalAmount(id int) (float32, error) {
	var err error = nil
	var total float32 = 0
	var items []model.InvoiceItem

	res := r.DB.Where("id_invoice = ?", id).Find(&items)
	if res.RowsAffected < 1 {
		log.Printf("not found invoice items")
		err = fmt.Errorf("not found invoice items")
	}

	for i := 0; i < len(items); i++ {
		total += items[i].Subtotal
	}

	return total, err
}

func NewMysqlPaymentGatewayRepository(db *gorm.DB) adapter.AdapterPaymentGatewayRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
