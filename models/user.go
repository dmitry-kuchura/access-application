package models

import (
	//"golang.org/x/crypto/bcrypt"
	"github.com/dmitry-kuchura/access-application/db"
	"strconv"
)

const insertUser = `
	INSERT INTO users (email, token, name, role)
	VALUES(?, ?, ?, 0) ON DUPLICATE KEY UPDATE
	token=VALUES(token), name=VALUES(name)
`

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

func GetUser(email, password string) (*User) {
	user := &User{}

	err := app.db.QueryRow("SELECT `id`, `name`, `token`, `email`, `password` FROM `users` WHERE `email` LIKE ?", email).Scan(
		&user.ID, &user.Name, &user.Token, &user.Email, &user.Password)

	if ValidatePassword(user.Password, password) && err != nil {
		return nil
	} else {
		return user
	}
}

func CreateUser(email, password, name string) (string, error) {
	res, err := Exec(insertUser, email, password, name, app.rand.String(10))

	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(id, 10), nil
}

func ValidatePassword(userPassword, password string) bool {

	//hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return true
}
