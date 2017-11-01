package models

import (
	"../app"
	"fmt"
)

const (
	insertFtp = `
	INSERT INTO ftp (domain_id, hostname, username, password, status, created_at, updated_at)
	VALUES(?, ?, ?, ?, 1, NOW(), NOW())
	`

	deleteFtp = `
	DELETE FROM ftp WHERE id = ?
	`

	selectFtp = `
	SELECT id, hostname, username, password, status FROM ftp WHERE domain_id = ?
	`
)

type Ftp struct {
	ID        int    `form:"id" json:"id"`
	DomainID  int    `form:"domain_id" json:"domain_id"`
	Hostname  string `form:"hostname" json:"hostname"`
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
	Status    int    `form:"status" json:"status"`
	CreatedAt uint8  `form:"created_at" json:"created_at"`
	UpdatedAt uint8  `form:"updated_at" json:"updated_at"`
}

func CreateFtp(domain int, hostname, username, password string) (bool, error) {
	_, err := app.Exec(insertFtp, domain, hostname, username, password)

	fmt.Println(err)

	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func SelectFtps(domain int) (ftps []Ftp, err error) {
	rows, err := app.Query(selectFtp, domain)

	for rows.Next() {
		d := Ftp{}
		err = rows.Scan(&d.ID, &d.Hostname, &d.Username, &d.Password, &d.Status)
		if err != nil {
			return ftps, err
		}
		ftps = append(ftps, d)
	}
	err = rows.Err()

	if err != nil {
		return ftps, err
	} else {
		return ftps, err
	}
}

func DeleteFtp(id int) (bool, error) {
	_, err := app.Exec(deleteFtp, id)

	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}
