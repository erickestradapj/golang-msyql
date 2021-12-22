package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// user:password@tcp/dbname
const url = "root:123456@tcp(localhost:3306)/goweb_db"

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	fmt.Println("=> Connection successfully")
	db = connection
}

func Close() {
	db.Close()
}
