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
	SELECT id, first_name, token, email, password FROM users WHERE status = 1 AND email LIKE ?
	`

	insertUser = `
	INSERT INTO users (email, password, first_name, token, role)
	VALUES(?, ?, ?, ?, 0) ON DUPLICATE KEY UPDATE
	token=VALUES(token), name=VALUES(name)
	`

	updateUser = `
	UPDATE users SET email = ?, first_name = ?, second_name = ?, updated_at = NOW() WHERE id = ?
	`

	changeStatusUser = `
	UPDATE  users SET status = CASE WHEN status = 0 THEN 1 ELSE 0 END WHERE id = ?
	`

	findUserByToken = `
	SELECT id, first_name, token, email, password FROM users WHERE status = 1 AND token LIKE ?
	`

	findUserByID = `
	SELECT id, first_name, token, email, password, status FROM users WHERE id = ?
	`

	deleteUser = `
	DELETE FROM users WHERE id = ?
	`

	selectAllUsers = `
	SELECT id, first_name, email, token, status, role
	FROM users
	LIMIT ?
	OFFSET ?
	`
	countAllUsers = `
	SELECT COUNT(*) as count
	FROM users
	`

	selectAllByFilter = `
	SELECT id, first_name, email, token, status, role
	FROM users
	WHERE status = ?
	`
)

type Identity interface {
	GetID() int
	GetName() string
}

type User struct {
	ID         int    `form:"id" json:"id"`
	FirstName  string `form:"first_name" json:"first_name"`
	SecondName string `form:"second_name" json:"second_name"`
	Email      string `form:"email" json:"email"`
	Password   string `form:"password" json:"password"`
	Token      string `form:"token" json:"token"`
	Status     int    `form:"status" json:"status"`
	Role       int    `form:"role" json:"role"`
}

func (u User) GetID() int {
	return u.ID
}

func (u User) GetName() string {
	return u.FirstName
}

// Поиск пользователя по Email
func GetUser(email, password string) (*User, bool) {
	user := &User{}

	err := app.QueryRow(selectUser, email).Scan(
		&user.ID, &user.FirstName, &user.Token, &user.Email, &user.Password)

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
		&user.ID, &user.FirstName, &user.Token, &user.Email, &user.Password,
	)

	if err == nil {
		return user, false
	} else {
		return nil, true
	}
}

// Обновление пользователя
func UpdateUser(email, firstName, secondName string, id int) error {
	_, err := app.Exec(updateUser, email, firstName, secondName, id)

	return err
}

// Создание пользователя
func CreateUser(email, password, firstName string) (string, error) {
	res, err := app.Exec(insertUser, email, hashedPassword(password), firstName, app.String(35))

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
			&user.ID, &user.FirstName, &user.Token, &user.Email, &user.Password, &user.Status)

		if err == nil {
			return user, false
		} else {
			return nil, true
		}
	} else {
		return nil, true
	}

}

// Получение списка пользователей с пагинацией
func AllUsers(param string) (users []User, count int, err error) {
	limit := 15
	page, _ := strconv.Atoi(param)
	offset := (page - 1) * limit

	rows, err := app.Query(selectAllUsers, limit, offset)

	row, _ := app.Query(countAllUsers)

	allUsers := app.CountRows(row)

	pages := allUsers / limit

	if err != nil {
		return users, pages, err
	}
	defer rows.Close()
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.FirstName, &u.Email, &u.Token, &u.Status, &u.Role)
		if err != nil {
			return users, pages, err
		}
		users = append(users, u)
	}
	err = rows.Err()
	return users, pages, err
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
