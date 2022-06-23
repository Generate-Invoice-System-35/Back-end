package model

import "time"

type User struct {
	ID         int       `json:"id" form:"id"`
	Username   string    `json:"username" form:"username"`
	Password   string    `json:"password" form:"password"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	Updated_At time.Time `json:"updated_at" form:"updated_at"`
}

func (u *User) TableName() string {
	// custom table name, this is default
	return "users"
}
