package model

import "time"

type User struct {
	ID         int        `json:"id" form:"id"`
	ID_Card    int        `json:"id_card" form:"id_card"`
	ID_Role    int        `json:"id_role" form:"id_role"`
	Username   string     `json:"username" form:"username"`
	Email      string     `json:"email" form:"email"`
	Password   string     `json:"password" form:"password"`
	Name       string     `json:"name" form:"name"`
	Born       *time.Time `json:"born" form:"born"`
	Phone      int        `json:"phone" form:"phone"`
	Street     string     `json:"street" form:"street"`
	City       string     `json:"city" form:"city"`
	State      string     `json:"state" form:"state"`
	Country    string     `json:"country" form:"country"`
	Zip        int        `json:"zip" form:"zip"`
	Created_At time.Time  `json:"created_at" form:"created_at"`
	Updated_At time.Time  `json:"updated_at" form:"updated_at"`
}
