package models

import (
	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model
		Name     string `json:"name"`
		Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
		Password string
	}

	TransformedUser struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
