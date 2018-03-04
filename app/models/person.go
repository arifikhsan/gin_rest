package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	Person struct {
		gorm.Model
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	TransformedPerson struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		Age       int       `json:"age"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
