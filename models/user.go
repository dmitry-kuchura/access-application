package models

import (
	"fmt"
	"strconv"

	"crypto/md5"
	"encoding/hex"

	"github.com/dmitry-kuchura/access-application/app"
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

	changeStatusUser = `
	UPDATE  users SET status = CASE WHEN status = 0 THEN 1 ELSE 0 END WHERE id = ?
	`

	findUserByToken = `
	SELECT id, name, token, email, password FROM users WHERE status = 1 AND token LIKE ?
	`

	findUserByID = `
	SELECT id, name, token, email, password, status FROM users WHERE id = ?
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

// Поиск пользователя по Email
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

// Поиск пользователя по Token
func FindUserByToken(token string) (*User, bool) {
	user := &User{}

	err := app.QueryRow(findUserByToken, token).Scan(
		&user.ID, &user.Name, &user.Token, &user.Email, &user.Password)

	if err == nil {
		return user, false
	} else {
		return nil, true
	}
}

// Создание пользователя
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

// Удаление пользователя
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

// Смена статуса пользователя и возврат текущего статуса
func ChangeStatusUser(id int) (*User, bool) {
	user := &User{}
	_, fail := app.Exec(changeStatusUser, id)

	if fail == nil {
		err := app.QueryRow(findUserByID, id).Scan(
			&user.ID, &user.Name, &user.Token, &user.Email, &user.Password, &user.Status)

		if err == nil {
			return user, false
		} else {
			return nil, true
		}
	} else {
		return nil, true
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
