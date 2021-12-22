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
	// user := models.CreateUser("Alexander", "456", "alexander@email.com")
	// fmt.Println(user)

	// users := models.ListUsers()
	// fmt.Println(users)

	user := models.GetUser(1)
	fmt.Println(user)

	// db.TruncateTable("users")
	db.Close()
	// db.Ping()
}
