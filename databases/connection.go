package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {

	var err error
	DB, err = gorm.Open("mysql", "root:toor@/demo?parseTime=true")

	if err != nil {
		panic("failed to connect database")
	}
	// defer DB.Close()
}
