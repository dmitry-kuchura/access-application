package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db, _ = sql.Open("mysql", "root:@/golang")

var Exec = db.Exec
var Query = db.Query
var QueryRow = db.QueryRow

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
