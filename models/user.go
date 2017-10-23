package models

import (
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"github.com/dmitry-kuchura/access-application/app"
	"fmt"
)

const insertUser = `
	INSERT INTO users (email, password, name, token, role)
	VALUES(?, ?, ?, ?, 0) ON DUPLICATE KEY UPDATE
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

	err := app.QueryRow("SELECT `id`, `name`, `token`, `email`, `password` FROM `users` WHERE `email` LIKE ?", email).Scan(
		&user.ID, &user.Name, &user.Token, &user.Email, &user.Password)

	fmt.Println(validPassword(password, user.Password))

	if validPassword(password, user.Password) && err != nil {
		return user
	} else {
		return nil
	}
}

func CreateUser(email, password, name string) (string, error) {
	res, err := app.Exec(insertUser, email, hashedPassword(password), name, app.String(25))

	if err == nil {
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(id, 10), nil
}

func hashedPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	return hashedPassword
}

func validPassword(password, currentPassword string) bool {
	myPassword := hashedPassword(password)


	fmt.Println(myPassword)
	fmt.Println(currentPassword)

	if myPassword == currentPassword {
		return true
	} else {
		return false
	}
}
