package models

import "../app"

const (
	selectMysql = `
	SELECT id, hostname, username, password, status FROM ftp WHERE domain_id = ?
	`
)

type Database struct {
	ID       int    `form:"id" json:"id"`
	DomainID int    `form:"domain_id" json:"domain_id"`
	Hostname string `form:"hostname" json:"hostname"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Status   int    `form:"status" json:"status"`
}

func SelectDatabases(domain int) (databases []Database, err error) {
	rows, err := app.Query(selectMysql, domain)

	for rows.Next() {
		d := Database{}
		err = rows.Scan(&d.ID, &d.Hostname, &d.Username, &d.Password, &d.Status)
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
