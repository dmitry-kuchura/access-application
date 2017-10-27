package models

import (
	"fmt"
	"strconv"

	"crypto/md5"
	"encoding/hex"

	"../app"
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

	selectAll = `
	SELECT id, name, email, token, status, role
	FROM users
	LIMIT ?
	OFFSET ?
	`
	countAll = `
	SELECT COUNT(*) as count
	FROM users
	`

	selectAllByFilter = `
	SELECT id, name, email, token, status, role
	FROM users
	WHERE status = ?
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

// Вторая и более страница
func AllUsers(param string) (users []User, count int, err error) {
	limit := 1
	page, _ := strconv.Atoi(param)
	offset := (page - 1) * limit

	rows, err := app.Query(selectAll, limit, offset)

	row, _ := app.Query(countAll)

	all := app.CountRows(row)

	fmt.Println(all)

	if err != nil {
		return users, all, err
	}
	defer rows.Close()
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Token, &u.Status, &u.Role)
		if err != nil {
			return users, all, err
		}
		users = append(users, u)
	}
	err = rows.Err()
	return users, all, err
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
