package mysql_test

import (
	"regexp"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	repository "Back-end/internal/payment_gateway/xendit/repository/mysql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// func TestCreateTransactionRecord(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	dial := mysql.Dialector{&mysql.Config{
// 		Conn:                      dbMock,
// 		SkipInitializeWithVersion: true,
// 	}}
// 	db, _ := gorm.Open(dial)
// 	repo := repository.NewMysqlPaymentGatewayRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectBegin()
// 	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
// 		WithArgs(1, "", "123451", "12345678901", time.Now(), time.Now()).
// 		WillReturnResult(driver.RowsAffected(1))
// 	fMock.ExpectCommit()

// 	err := repo.CreateInvoice(model.Invoice{
// 		ID_Payment_Status: 1,
// 		Number:            "INV/2020/0001",
// 		Name:              "issuer name test",
// 		Email:             "testgmail@gmail.com",
// 		Phone_Number:      "12315415",
// 		Address:           "address test",
// 		Description:       "-",
// 		Invoice_Date:      time.Now(),
// 		Due_Date:          time.Now(),
// 		Created_At:        time.Now(),
// 		Updated_At:        time.Now(),
// 		ID:                1,
// 	})
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }

func TestGetIDInvoicePayment(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlPaymentGatewayRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `transaction_records`")).
		WillReturnRows(sqlmock.NewRows([]string{"id_invoice", "id_invoice_payment", "id_user_payment"}).
			AddRow(1, "12345671", "1234567891").
			AddRow(2, "12345672", "1234567892"))

	res, err := repo.GetIDInvoicePayment(1)
	assert.Equal(t, res.ID_Invoice_Payment, "12345671")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetInvoices(t *testing.T) {

}

func TestGetTotalAmount(t *testing.T) {

}

func TestUpdateStatusInvoice(t *testing.T) {

}
