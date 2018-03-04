package migrations

import (
	"fmt"

	"github.com/arifikhsan/gin_rest/app/models"
	databases "github.com/arifikhsan/gin_rest/databases"
)

func init() {
	if databases.DB.HasTable(&models.Person{}) {
		fmt.Println("Person table has been created")
	}
	databases.DB.AutoMigrate(&models.Person{})
	fmt.Println("Updating person table...")

}
