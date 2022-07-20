package repository_test

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Back-end/internal/upload_file/model"
	"Back-end/internal/upload_file/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateImage(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := repository.NewMysqlUploadRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs("Image Testing", "main-qimg-55949c34563b674c37450e20ce61fa9b-lq.jpg", 12352, time.Now(), time.Now(), 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.CreateImage(model.File{
		Name:       "Image Testing",
		File_Name:  "main-qimg-55949c34563b674c37450e20ce61fa9b-lq.jpg",
		File_Size:  12352,
		Created_At: time.Now(),
		Updated_At: time.Now(),
		ID:         1,
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

// func TestGetAllImages(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	dial := mysql.Dialector{&mysql.Config{
// 		Conn:                      dbMock,
// 		SkipInitializeWithVersion: true,
// 	}}
// 	db, _ := gorm.Open(dial)
// 	repo := repository.NewMysqlUploadRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `files`")).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "number", "file_name", "file_size"}).
// 			AddRow(1, "Image Testing 1", "main-qimg-55949c34563b674c37450e20ce61fa9b-lq1.jpg", 123521).
// 			AddRow(2, "Image Testing 2", "main-qimg-55949c34563b674c37450e20ce61fa9b-lq2.jpg", 123522))

// 	res := repo.GetAllImages()
// 	assert.Equal(t, res[0].Name, "Image Testing 1")
// 	assert.Len(t, res, 2)
// }

// func TestGetImageByID(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	dial := mysql.Dialector{&mysql.Config{
// 		Conn:                      dbMock,
// 		SkipInitializeWithVersion: true,
// 	}}
// 	db, _ := gorm.Open(dial)
// 	repo := repository.NewMysqlUploadRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `files`")).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "number", "file_name", "file_size"}).
// 			AddRow(1, "Image Testing 1", "main-qimg-55949c34563b674c37450e20ce61fa9b-lq1.jpg", 123521).
// 			AddRow(2, "Image Testing 2", "main-qimg-55949c34563b674c37450e20ce61fa9b-lq2.jpg", 123522))

// 	res, err := repo.GetImageByID(1)
// 	assert.Equal(t, res.Name, "Image Testing 1")
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }

// func TestUpdateImage(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	dial := mysql.Dialector{&mysql.Config{
// 		Conn:                      dbMock,
// 		SkipInitializeWithVersion: true,
// 	}}
// 	db, _ := gorm.Open(dial)
// 	repo := repository.NewMysqlUploadRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectBegin()
// 	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
// 		WithArgs("Image Testing Update", 1).
// 		WillReturnResult(sqlmock.NewResult(0, 1))
// 	fMock.ExpectCommit()

// 	err := repo.UpdateImage(1, model.File{
// 		Name: "Image Testing Update",
// 	})
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }

// func TestDeleteImage(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	dial := mysql.Dialector{&mysql.Config{
// 		Conn:                      dbMock,
// 		SkipInitializeWithVersion: true,
// 	}}
// 	db, _ := gorm.Open(dial)
// 	repo := repository.NewMysqlUploadRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectBegin()
// 	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
// 		WithArgs(1).
// 		WillReturnResult(driver.RowsAffected(1))
// 	fMock.ExpectCommit()

// 	err := repo.DeleteImage(1)
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }
