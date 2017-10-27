package models

import (
	"strconv"
	"errors"
	"../app"
)

const (
	insertDomain = `
	INSERT INTO domains (name, url, status)
	VALUES(?, ?, 1) ON DUPLICATE KEY UPDATE
	name = VALUES(name)
	`

	checkDomain = `
	SELECT COUNT(*) as count FROM domains WHERE name LIKE ? OR url LIKE ?
	`
)

type Domains struct {
	ID     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`
	Url    string `form:"url" json:"url"`
	Status int    `form:"status" json:"status"`
}

func CreateDomain(name, url string) (string, error) {
	if CheckDomain(name, url) {
		res, err := app.Exec(insertDomain, name, url)

		id, err := res.LastInsertId()
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(id, 10), nil
	} else {
		err := errors.New("Domain was already created!")
		return "", err
	}
}

func CheckDomain(name, url string) bool {
	res, _ := app.Query(checkDomain, name, url)

	if app.CountRows(res) >= 1 {
		return false
	} else {
		return true
	}
}
