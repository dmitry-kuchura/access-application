package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const insertUser = `
	INSERT INTO users (email, password, name, token, role)
	VALUES(?, ?, ?, ?, 0) ON DUPLICATE KEY UPDATE
	token=VALUES(token), name=VALUES(name)
`

var db, _ = sql.Open("mysql", "root:@/golang")

var Exec = db.Exec
var Query = db.Query
var QueryRow = db.QueryRow