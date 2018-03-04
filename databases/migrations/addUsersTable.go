package migrations

import (
	"fmt"

	"github.com/arifikhsan/gin_rest/app/models"
	databases "github.com/arifikhsan/gin_rest/databases"
)

func init() {
	if databases.DB.HasTable(&models.User{}) {
		fmt.Println("User table has been created")
	}
	databases.DB.AutoMigrate(&models.User{})
	fmt.Println("Updating user table...")

}
