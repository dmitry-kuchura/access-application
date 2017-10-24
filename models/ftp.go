package models

type Ftp struct {
	ID       int    `form:"id" json:"id"`
	DomainID int    `form:"domain_id" json:"domain_id"`
	Hostname string `form:"hostname" json:"hostname"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Status   int    `form:"status" json:"status"`
}
