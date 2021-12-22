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

	// user := models.GetUser(1)
	// fmt.Println(user)

	// user := models.GetUser(2)
	// fmt.Println(user)
	// user.Username = "alexander"
	// user.Password = "123"
	// user.Email = "a@email.com"
	// user.Save()

	// user := models.GetUser(2)
	// fmt.Println(user)
	// user.Delete()

	user := models.CreateUser("Alexander", "456", "alexander@email.com")
	fmt.Println(user)

	// db.TruncateTable("users")
	// fmt.Println(models.ListUsers())

	db.Close()
	// db.Ping()
}
