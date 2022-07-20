package repository_test

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Back-end/internal/invoice/model"
	"Back-end/internal/invoice/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlInvoiceRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs(1, "INV/2020/0001", "issuer name test", "testgmail@gmail.com", "12315415", "address test", "-", time.Now(), time.Now(), time.Now(), time.Now(), 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.CreateInvoice(model.Invoice{
		ID_Payment_Status: 1,
		Number:            "INV/2020/0001",
		Name:              "issuer name test",
		Email:             "testgmail@gmail.com",
		Phone_Number:      "12315415",
		Address:           "address test",
		Description:       "-",
		Invoice_Date:      time.Now(),
		Due_Date:          time.Now(),
		Created_At:        time.Now(),
		Updated_At:        time.Now(),
		ID:                1,
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAllInvoices(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlInvoiceRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `invoices`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_payment_status", "number", "name", "email", "phone_number", "address", "description"}).
			AddRow(1, 1, "INV/2022/0001", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1", "1-1").
			AddRow(2, 1, "INV/2022/0002", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2", "2-2"))

	res := repo.GetAllInvoices()
	assert.Equal(t, res[0].Number, "INV/2022/0001")
	assert.Len(t, res, 2)
}

func TestGetInvoiceByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlInvoiceRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `invoices`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_payment_status", "number", "name", "email", "phone_number", "address", "description"}).
			AddRow(1, 1, "INV/2022/0001", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1", "1-1").
			AddRow(2, 1, "INV/2022/0002", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2", "2-2"))

	res, err := repo.GetInvoiceByID(1)
	assert.Equal(t, res.Number, "INV/2022/0001")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetInvoicesPagination(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlInvoiceRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `invoices`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_payment_status", "number", "name", "email", "phone_number", "address", "description"}).
			AddRow(1, 1, "INV/2022/0001", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1", "1-1").
			AddRow(2, 1, "INV/2022/0002", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2", "2-2"))

	data := model.Pagination{
		Page: 1,
	}
	res, err := repo.GetInvoicesPagination(data)
	assert.Equal(t, res[0].Number, "INV/2022/0001")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetInvoicesByPaymentStatus(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlInvoiceRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `invoices`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_payment_status", "number", "name", "email", "phone_number", "address", "description"}).
			AddRow(1, 1, "INV/2022/0001", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1", "1-1").
			AddRow(2, 1, "INV/2022/0002", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2", "2-2"))

	res, err := repo.GetInvoicesByPaymentStatus(1)
	assert.Equal(t, res[0].Number, "INV/2022/0001")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetInvoicesByNameCustomer(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlInvoiceRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `invoices`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_payment_status", "number", "name", "email", "phone_number", "address", "description"}).
			AddRow(1, 1, "INV/2022/0001", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1", "1-1").
			AddRow(2, 1, "INV/2022/0002", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2", "2-2"))

	res, err := repo.GetInvoicesByNameCustomer("name testing 1")
	assert.Equal(t, res[0].Number, "INV/2022/0001")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateInvoiceByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlInvoiceRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("INV/2022/0030", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateInvoiceByID(1, model.Invoice{
		Number: "INV/2022/0030",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteInvoiceByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlInvoiceRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteInvoiceByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}
