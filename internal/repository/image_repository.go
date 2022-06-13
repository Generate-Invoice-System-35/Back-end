package repository

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

func (r *RepositoryMysqlLayer) CreateImage(image model.File) error {
	res := r.DB.Create(&image)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllImages() []model.File {
	images := []model.File{}
	r.DB.Find(&images)

	return images
}

func (r *RepositoryMysqlLayer) GetImageByID(id int) (model.File, string, error) {
	var image model.File

	res := r.DB.Where("id = ?", id).Find(&image)
	if res.RowsAffected < 1 {
		return image, "", fmt.Errorf("not found")
	}

	filePath := "storage/"
	urlImage := filePath + image.File_Name

	return image, urlImage, nil
}

func (r *RepositoryMysqlLayer) UpdateImage(id int, image model.File) error {
	// Finding data
	var img model.File
	res1 := r.DB.Where("id = ?", id).Find(&img)
	if res1.RowsAffected < 1 {
		return fmt.Errorf("error find image")
	}

	// Delete file in local storage
	filePath := "storage/"
	err := os.Remove(filePath + img.File_Name)
	if err != nil {
		log.Fatal(err)
	}

	// Update data in database
	res2 := r.DB.Where("id = ?", id).UpdateColumns(&image)
	if res2.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteImage(id int) error {
	// Finding data
	var img model.File
	res1 := r.DB.Where("id = ?", id).Find(&img)
	if res1.RowsAffected < 1 {
		return fmt.Errorf("error find image")
	}

	// Delete file in local storage
	filePath := "storage/"
	err := os.Remove(filePath + img.File_Name)
	if err != nil {
		log.Fatal(err)
	}

	// Delete data in database
	res2 := r.DB.Delete(&model.File{ID: id})
	if res2.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlUploadRepository(db *gorm.DB) adapter.AdapterUploadImageRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
