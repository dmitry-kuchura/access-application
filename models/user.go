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

func GetToken(UserEmail string) {
	stmt, _ := db.Prepare("INSERT `users` SET email=?,pasword=?,token=?,name=?")
	res, _ := stmt.Exec("test@domail.com", "Swqa123123123", "sa6d7587d6df7gnh5f2jm5fjm", "Andrew")
	id, _ := res.LastInsertId()

	fmt.Println(id)
	fmt.Println(UserEmail)

	//db.QueryRow("SELECT `id`, `email`, `password`, `token`, `name` FROM users WHERE email = $1", UserEmail)

	//user := new(User)
	//result := row.Scan(&user.ID, &user.Email, &user.Password, &user.Token, &user.Name)
	//
	//fmt.Println(result)
}
