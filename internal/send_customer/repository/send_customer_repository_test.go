package repository_test

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Back-end/internal/send_customer/model"
	"Back-end/internal/send_customer/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlSendCustomerRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs("sendcustomer@gmail.com", "Test Subject Email", "Test Body Email", time.Now(), time.Now(), 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.SendEmail(model.SendCustomer{
		To:         "sendcustomer@gmail.com",
		Subject:    "Test Subject Email",
		Body:       "Test Body Email",
		Created_At: time.Now(),
		Updated_At: time.Now(),
		ID:         1,
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
