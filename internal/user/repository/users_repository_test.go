package repository_test

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Back-end/internal/user/model"
	"Back-end/internal/user/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "name", "email", "phone_number", "address"}).
			AddRow(1, "usernametesting1", "passwordtesting1", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1").
			AddRow(2, "usernametesting2", "passwordtesting2", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2"))

	res := repo.GetAllUsers()
	assert.Equal(t, res[0].Username, "usernametesting1")
	assert.Len(t, res, 2)
}

func TestGetUserByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "name", "email", "phone_number", "address"}).
			AddRow(1, "usernametesting1", "passwordtesting1", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1").
			AddRow(2, "usernametesting2", "passwordtesting2", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2"))

	res, err := repo.GetUserByID(1)
	assert.Equal(t, res.Username, "usernametesting1")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUsernameExist(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "name", "email", "phone_number", "address"}).
			AddRow(1, "usernametesting1", "passwordtesting1", "name testing 1", "emailtesting1@gmail.com", "081234567891", "address testing 1").
			AddRow(2, "usernametesting2", "passwordtesting2", "name testing 2", "emailtesting2@gmail.com", "081234567892", "address testing 2"))

	res, err := repo.UsernameExist("usernametesting1")
	assert.Equal(t, res.Username, "usernametesting1")
	assert.Error(t, err)
	assert.True(t, true)
}

func TestUpdateUserByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("usernametesting", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateUserByID(1, model.User{
		Username: "usernametesting",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteUserByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteUserByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}
