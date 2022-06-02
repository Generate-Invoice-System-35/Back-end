package repository_test

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"Back-end/internal/model"
	"Back-end/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		WithArgs("maderahano", "123456", time.Now(), time.Now(), 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.Register(model.User{
		Username:   "maderahano",
		Password:   "123456",
		Created_At: time.Now(),
		Updated_At: time.Now(),
		ID:         1,
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func UsernameExists(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlAuthRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
			AddRow(1, "maderahano", "123456"))

	res, err := repo.UsernameExists("maderahano")
	assert.Equal(t, res.Username, "maderahano")
	assert.Error(t, err)
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
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
			AddRow(1, "maderahano", "123456"))

	res, err := repo.Login("maderahano")
	assert.Equal(t, res.Username, "maderahano")
	assert.NoError(t, err)
	assert.True(t, true)
}
