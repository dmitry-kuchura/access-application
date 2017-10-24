package models

import (
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"github.com/dmitry-kuchura/access-application/app"
	"fmt"
)

const (
	selectUser = `
	SELECT id, name, token, email, password FROM users WHERE status = 1 AND email LIKE ?
	`

	insertUser = `
	INSERT INTO users (email, password, name, token, role)
	VALUES(?, ?, ?, ?, 0) ON DUPLICATE KEY UPDATE
	token=VALUES(token), name=VALUES(name)
	`

	findUser = `
	SELECT id, name, token, email, password FROM users WHERE status = 1 AND token LIKE ?
	`

	deleteUser = `
	DELETE FROM users WHERE id = ?
	`
)

type Identity interface {
	GetID() int
	GetName() string
}

type User struct {
	ID       int    `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Token    string `form:"token" json:"token"`
	Status   int    `form:"status" json:"status"`
	Role     int    `form:"role" json:"role"`
}

func (u User) GetID() int {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func GetUser(email, password string) (*User, bool) {
	user := &User{}

	err := app.QueryRow(selectUser, email).Scan(
		&user.ID, &user.Name, &user.Token, &user.Email, &user.Password)

	if validPassword(password, user.Password) && err == nil {
		return user, false
	} else {
		return nil, true
	}
}

func FindUserByToken(token string) (*User, bool) {
	user := &User{}

	err := app.QueryRow(findUser, token).Scan(
		&user.ID, &user.Name, &user.Token, &user.Email, &user.Password)

	if err == nil {
		return user, false
	} else {
		return nil, true
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

func DeleteUser(id int) (bool, error) {
	_, err := app.Exec(deleteUser, id)

	fmt.Println(err)
	fmt.Println(id)

	if err == nil {
		return true, nil
	} else {
		return false, err
	}
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
