package models

import "github.com/dmitry-kuchura/access-application/app"

const (
	selectMysql = `
	SELECT id, domain_id, phpmyadmin, hostname, username, password, status, created_at, updated_at FROM mysql WHERE domain_id = ?
	`
)

type Database struct {
	ID        int    `form:"id" json:"id"`
	DomainID  int    `form:"domain_id" json:"domain_id"`
	MyAdmin   string `form:"phpmyadmin" json:"phpmyadmin"`
	Hostname  string `form:"hostname" json:"hostname"`
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
	Status    int    `form:"status" json:"status"`
	CreatedAt string `form:"created_at" json:"created_at"`
	UpdatedAt string `form:"updated_at" json:"updated_at"`
}

func SelectDatabases(domain int) (databases []Database, err error) {
	rows, err := app.Query(selectMysql, domain)

	for rows.Next() {
		d := Database{}
		err = rows.Scan(&d.ID, &d.DomainID, &d.MyAdmin, &d.Hostname, &d.Username, &d.Password, &d.Status, &d.CreatedAt, &d.UpdatedAt)
		if err != nil {
			return databases, err
		}
		databases = append(databases, d)
	}
	err = rows.Err()

	if err != nil {
		return databases, err
	} else {
		return databases, err
	}
}
