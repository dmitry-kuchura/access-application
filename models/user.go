package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB
var err error

type Identity interface {
	GetID() int
	GetName() string
}

type User struct {
	ID       int    `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Token    string `form:"token" json:"token"`
}

func (u User) GetID() int {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func GetToken(email string) {
	var user User
	
	row := db.QueryRow("SELECT `id`, `email`, `password`, `token`, `name` FROM `users` WHERE `id` = $1", email)
	row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Token)

	return user
}
