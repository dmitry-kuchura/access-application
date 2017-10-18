package models

type Identity interface {
	GetID() int
	GetName() string
	GetToken() string
}

type User struct {
	ID       int `form:"id" json:"id"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Token    string `form:"token" json:"token"`
	Name     string `form:"name" json:"name"`
}

func (u User) GetID() int {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetToken() string {
	return u.Token
}
