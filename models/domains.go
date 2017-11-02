package models

import (
	"strconv"
	"errors"

	"github.com/dmitry-kuchura/access-application/app"
)

const (
	insertDomain = `
	INSERT INTO domains (name, url, description, status, created_at, updated_at)
	VALUES(?, ?, ?, 1, NOW(), NOW()) ON DUPLICATE KEY UPDATE
	name = VALUES(name)
	`

	checkDomain = `
	SELECT COUNT(*) as count FROM domains WHERE name LIKE ? OR url LIKE ?
	`

	selectAllDomains = `
	SELECT id, name, url, status, updated_at
	FROM domains
	LIMIT ?
	OFFSET ?
	`
	countAllDomains = `
	SELECT COUNT(*) as count
	FROM domains
	`

	deleteDomain = `
	DELETE FROM domains WHERE id = ?
	`

	selectDomain = `
	SELECT id, name, url, status, updated_at FROM domains WHERE id = ?
	`
)

type Domain struct {
	ID          int        `form:"id" json:"id"`
	Name        string     `form:"name" json:"name"`
	Url         string     `form:"url" json:"url"`
	Description string     `form:"description" json:"description"`
	Status      int        `form:"status" json:"status"`
	Updated     string     `form:"updated_at" json:"updated_at"`
	Ftps        []Ftp      `form:"ftp" json:"ftp"`
	Databases   []Database `form:"database" json:"database"`
	Hostings    []Hosting  `form:"hosting" json:"hosting"`
	Admins      []Admin    `form:"admin" json:"admin"`
}

// Добавление домена
func CreateDomain(name, url, description string) (string, error) {
	if CheckDomain(name, url) {
		res, err := app.Exec(insertDomain, name, url, description)

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

// Список доменов
func AllDomains(param string) (domains []Domain, count int, err error) {
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
		d := Domain{}
		err = rows.Scan(&d.ID, &d.Name, &d.Url, &d.Status, &d.Updated)
		if err != nil {
			return domains, pages, err
		}
		domains = append(domains, d)
	}
	err = rows.Err()
	return domains, pages, err
}

// Удаление домена
func DeleteDomain(id int) (bool, error) {
	_, err := app.Exec(deleteDomain, id)

	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

// Получение конкретного домена
func GetDomain(param int) (domains []Domain, err error) {
	row, err := app.Query(selectDomain, param)
	defer row.Close()

	ftp, err := SelectFtps(param)
	mysql, err := SelectDatabases(param)
	admin, err := SelectAdmins(param)
	hosting, err := SelectHosting(param)

	for row.Next() {
		d := Domain{}
		err = row.Scan(&d.ID, &d.Name, &d.Url, &d.Status, &d.Updated)
		if err != nil {
			return domains, err
		}

		d.Ftps = ftp
		d.Databases = mysql
		d.Admins = admin
		d.Hostings = hosting

		domains = append(domains, d)
	}
	err = row.Err()

	return domains, err
}

func CheckDomain(name, url string) bool {
	res, _ := app.Query(checkDomain, name, url)

	if app.CountRows(res) >= 1 {
		return false
	} else {
		return true
	}
}
