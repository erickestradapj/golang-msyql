package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()

	// fmt.Println(db.ExistsTable("users"))
	// db.CreateTable(models.UserSchema, "users")
	user := models.CreateUser("Erick", "123", "erick@email.com")
	fmt.Println(user)
	// db.TruncateTable("users")

	db.Close()
	// db.Ping()
}
