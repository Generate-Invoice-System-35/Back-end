package repository_test

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Back-end/internal/auth/repository"
	"Back-end/internal/user/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlAuthRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs("usernametesting1", "passwordtesting1", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1", time.Now(), time.Now(), 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.Register(model.User{
		Username:     "usernametesting1",
		Password:     "passwordtesting1",
		Name:         "name testing 1",
		Email:        "emailtesting1@gmail.com",
		Phone_Number: "081234567891",
		Address:      "address testing 1",
		Created_At:   time.Now(),
		Updated_At:   time.Now(),
		ID:           1,
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestLogin(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlAuthRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "name", "email", "phone_number", "address"}).
			AddRow(1, "usernametesting1", "passwordtesting1", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1").
			AddRow(2, "usernametesting2", "passwordtesting2", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2"))

	res, err := repo.Login("usernametesting1")
	assert.Equal(t, res.Username, "usernametesting1")
	assert.NoError(t, err)
}

func TestUsernameExists(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlAuthRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "name", "email", "phone_number", "address"}).
			AddRow(1, "usernametesting1", "passwordtesting1", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1").
			AddRow(2, "usernametesting2", "passwordtesting2", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2"))

	res, err := repo.UsernameExists("usernametesting1")
	assert.Equal(t, res.Username, "usernametesting1")
	assert.Error(t, err)
	assert.True(t, true)
}
