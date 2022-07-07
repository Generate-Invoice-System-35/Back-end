package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Back-end/config"
	invoice "Back-end/internal/invoice/model"
	item "Back-end/internal/invoice_item/model"
	status "Back-end/internal/invoice_payment_status/model"
	transaction "Back-end/internal/payment_gateway/xendit/model"
	send "Back-end/internal/send_customer/model"
	file "Back-end/internal/upload_file/model"
	user "Back-end/internal/user/model"
)

func InitDB(conf config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err.Error())
	}

	initMigrate(DB)
	return DB
}

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&file.File{})
	db.AutoMigrate(&invoice.Invoice{})
	db.AutoMigrate(&item.InvoiceItem{})
	db.AutoMigrate(&send.SendCustomer{})
	db.AutoMigrate(&transaction.TransactionRecord{})
	db.AutoMigrate(&status.InvoicePaymentStatus{})
}
