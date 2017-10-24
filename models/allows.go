package models

const (
	selectAllows = `
	SELECT id, name, token, email, password FROM users WHERE status = 1 AND email LIKE ?
	`

	updateAllows = `
	INSERT INTO users (email, password, name, token, role)
	VALUES(?, ?, ?, ?, 0) ON DUPLICATE KEY UPDATE
	token=VALUES(token), name=VALUES(name)
	`
)

type Allows struct {
	ID       int    `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Token    string `form:"token" json:"token"`
	Status   int    `form:"status" json:"status"`
	Role     int    `form:"role" json:"role"`
}
