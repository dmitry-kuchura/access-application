package models

import (
	"fmt"
	"github.com/dmitry-kuchura/access-application/app"
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
	SELECT id, domain_id, hostname, username, password, status, created_at, updated_at FROM ftp WHERE domain_id = ?
	`
)

type Ftp struct {
	ID        int    `form:"id" json:"id"`
	DomainID  int    `form:"domain_id" json:"domain_id"`
	Hostname  string `form:"hostname" json:"hostname"`
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
	Status    int    `form:"status" json:"status"`
	CreatedAt string `form:"created_at" json:"created_at"`
	UpdatedAt string `form:"updated_at" json:"updated_at"`
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
		f := Ftp{}
		err = rows.Scan(&f.ID, &f.DomainID, &f.Hostname, &f.Username, &f.Password, &f.Status, &f.CreatedAt, &f.UpdatedAt)
		if err != nil {
			return ftps, err
		}
		ftps = append(ftps, f)
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
