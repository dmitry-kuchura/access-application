package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db, _ = sql.Open("mysql", "root:@/golang")
	Exec = db.Exec
	Query = db.Query
	QueryRow = db.QueryRow
)

func CountRows(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		CheckErr(err)
	}
	return count
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
