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

	selectAllDomains = `
	SELECT id, name, url, status
	FROM domains
	LIMIT ?
	OFFSET ?
	`
	countAllDomains = `
	SELECT COUNT(*) as count
	FROM domains
	`
)

type Domains struct {
	ID     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`
	Url    string `form:"url" json:"url"`
	Status int    `form:"status" json:"status"`
}

// Добавление домена
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

func AllDomains(param string) (domains []Domains, count int, err error)  {
	limit := 15
	page, _ := strconv.Atoi(param)
	offset := (page - 1) * limit

	rows, err := app.Query(selectAllDomains, limit, offset)

	row, _ := app.Query(countAllDomains)

	allDomain := app.CountRows(row)

	pages := allDomain / limit

	if err != nil {
		return domains, pages, err
	}
	defer rows.Close()
	for rows.Next() {
		d := Domains{}
		err = rows.Scan(&d.ID, &d.Name, &d.Url, &d.Status)
		if err != nil {
			return domains, pages, err
		}
		domains = append(domains, d)
	}
	err = rows.Err()
	return domains, pages, err
}


func CheckDomain(name, url string) bool {
	res, _ := app.Query(checkDomain, name, url)

	if app.CountRows(res) >= 1 {
		return false
	} else {
		return true
	}
}
