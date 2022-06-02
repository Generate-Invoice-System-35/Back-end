package repository_test

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"Back-end/internal/model"
	"Back-end/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
			AddRow(1, "maderahano", "123456").
			AddRow(2, "fkindarya", "123456").
			AddRow(3, "anovianto", "123456"))

	res := repo.GetAllUsers()
	assert.Equal(t, res[0].Username, "maderahano")
	assert.Equal(t, res[1].Username, "fkindarya")
	assert.Equal(t, res[2].Username, "anovianto")
	assert.Len(t, res, 3)
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
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
			AddRow(1, "maderahano", "123456"))

	res, err := repo.GetUserByID(1)
	assert.Equal(t, res.Username, "maderahano")
	assert.NoError(t, err)
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
		WithArgs("maderahano", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateUserByID(1, model.User{
		Username: "maderahano",
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
