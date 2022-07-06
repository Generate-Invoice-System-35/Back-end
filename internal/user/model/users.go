package model

import "time"

type User struct {
	ID           int       `json:"id" form:"id"`
	Username     string    `json:"username" form:"username"`
	Password     string    `json:"password" form:"password"`
	Name         string    `json:"name" form:"name"`
	Email        string    `json:"email" form:"email"`
	Phone_Number string    `json:"phone_number" form:"phone_number"`
	Address      string    `json:"address" form:"address"`
	Created_At   time.Time `json:"created_at" form:"created_at"`
	Updated_At   time.Time `json:"updated_at" form:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
