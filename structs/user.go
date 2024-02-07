package structs

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Age   uint   `json:"age,omitempty"`
}
