package models

import "github.com/dmitry-kuchura/access-application/app"

const (
	selectHosting = `
	SELECT id, domain_id, url, username, password, status, created_at, updated_at FROM hosting WHERE domain_id = ?
	`
)

type Hosting struct {
	ID        int    `form:"id" json:"id"`
	DomainID  int    `form:"domain_id" json:"domain_id"`
	Url       string `form:"url" json:"url"`
	Login     string `form:"login" json:"login"`
	Password  string `form:"password" json:"password"`
	Status    int    `form:"status" json:"status"`
	CreatedAt string `form:"created_at" json:"created_at"`
	UpdatedAt string `form:"updated_at" json:"updated_at"`
}

func SelectHosting(domain int) (hosting []Hosting, err error) {
	rows, err := app.Query(selectHosting, domain)

	for rows.Next() {
		h := Hosting{}
		err = rows.Scan(&h.ID, &h.DomainID, &h.Url, &h.Login, &h.Password, &h.Status, &h.CreatedAt, &h.UpdatedAt)
		if err != nil {
			return hosting, err
		}
		hosting = append(hosting, h)
	}
	err = rows.Err()

	if err != nil {
		return hosting, err
	} else {
		return hosting, err
	}
}
