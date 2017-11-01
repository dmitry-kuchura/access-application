package models

type Admin struct {
	ID        int    `form:"id" json:"id"`
	Name      string `form:"name" json:"name"`
	Url       string `form:"url" json:"url"`
	Status    int    `form:"status" json:"status"`
	CreatedAt string `form:"created_at" json:"created_at"`
	UpdatedAt string `form:"updated_at" json:"updated_at"`
}
