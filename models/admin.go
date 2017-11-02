package models

import "github.com/dmitry-kuchura/access-application/app"

const (
	insertAdmin = `
	INSERT INTO backend (domain_id, hostname, username, password, status, created_at, updated_at)
	VALUES(?, ?, ?, ?, 1, NOW(), NOW())
	`

	deleteAdmin = `
	DELETE FROM backend WHERE id = ?
	`

	selectAdmin = `
	SELECT id, domain_id, url, login, password, status, created_at, updated_at FROM backend WHERE domain_id = ?
	`
)

type Admin struct {
	ID        int    `form:"id" json:"id"`
	DomainID  int    `form:"domain_id" json:"domain_id"`
	Url       string `form:"url" json:"url"`
	Login     string `form:"login" json:"login"`
	Password  string `form:"password" json:"password"`
	Status    int    `form:"status" json:"status"`
	CreatedAt string `form:"created_at" json:"created_at"`
	UpdatedAt string `form:"updated_at" json:"updated_at"`
}

func SelectAdmins(domain int) (admin []Admin, err error) {
	rows, err := app.Query(selectAdmin, domain)

	for rows.Next() {
		f := Admin{}
		err = rows.Scan(&f.ID, &f.DomainID, &f.Url, &f.Login, &f.Password, &f.Status, &f.CreatedAt, &f.UpdatedAt)
		if err != nil {
			return admin, err
		}
		admin = append(admin, f)
	}
	err = rows.Err()

	if err != nil {
		return admin, err
	} else {
		return admin, err
	}
}