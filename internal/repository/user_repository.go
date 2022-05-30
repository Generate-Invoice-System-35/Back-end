package repository

import "gorm.io/gorm"

type RepositoryMysqlLayer struct {
	DB *gorm.DB
}
