package main

import (
	_ "github.com/arifikhsan/gin_rest/databases/migrations"
	_ "github.com/arifikhsan/gin_rest/routers"
)

func main() {
	// router.Run(":8080")
	// defer connection.DB.Close()
}
