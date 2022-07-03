package model

import "time"

type File struct {
	ID         int       `json:"id" form:"id"`
	Name       string    `json:"name" form:"name"`
	File_Name  string    `json:"file_name" form:"file_name"`
	File_Size  int       `json:"file_size" form:"file_size"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	Updated_At time.Time `json:"updated_at" form:"updated_at"`
}
