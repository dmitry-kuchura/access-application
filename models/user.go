package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/dmitry-kuchura/access-application/app"
)

var db, _ = sql.Open("mysql", app.Config.DSN)

var Exec = db.Exec
var Query = db.Query
var QueryRow = db.QueryRow

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

func GetUser(email string) (*User) {
	user := &User{}

	err := db.QueryRow("SELECT `id`, `name`, `token`, `email` FROM `users` WHERE `id`=?", 1).Scan(
		&user.ID, &user.Name, &user.Token, &user.Email)

	if err != nil {
		return nil
	} else {
		return user
	}
}
