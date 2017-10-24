package models

type Domains struct {
	ID     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`
	Url    string `form:"url" json:"url"`
	Status int    `form:"status" json:"status"`
}
