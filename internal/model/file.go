package model

type File struct {
	ID        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	File_Name string `json:"file_name" form:"file_name"`
	File_Size int    `json:"file_size" form:"file_size"`
}
