package models

import "gomysql/db"

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

type Users []User

/* ===== SCHEMA USER ===== */
const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

/* ===== BUILD USER ===== */
func NewUser(username, password, email string) *User {
	user := &User{
		Username: username,
		Password: password,
		Email:    email,
	}

	return user
}

/* ===== CREATE USER AND INSERT IN DB ===== */
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()

	return user
}

/* ===== INSERT ROW - method===== */
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

/* ===== LIST ALL ROWS===== */
func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}

/* ===== GET A ROW ===== */
func GetUser(id int) *User {
	user := NewUser("", "", "")

	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}

	return user
}

/* ===== UPDATE ROW===== */
func (user *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

/* ===== SAVE OR EDIT ROW ===== */
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

/* ===== DELETE ROW ===== */
func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}
