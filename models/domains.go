package models

import (
	//"strconv"
	"github.com/dmitry-kuchura/access-application/app"
	"fmt"
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

	CheckDomain(name, url)

	//res, err := app.Exec(insertDomain, name, url)
	//
	//id, err := res.LastInsertId()
	//if err != nil {
	//	return "", err
	//}
	//return strconv.FormatInt(id, 10), nil

	return "", nil
}

func CheckDomain(name, url string) bool {

	res, err := app.Query(checkDomain, name, url)

	fmt.Println("Total count:", app.checkCountRows(res))
	fmt.Println(err)

	return true
}
