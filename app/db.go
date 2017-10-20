package app

//import (
//	"database/sql"
//	"log"
//)
//
//type DB struct {
//	db  *sql.DB
//	dsn string
//	log *log.Logger
//}
//
//func NewDB(driver, dsn string, log *log.Logger) (*DB, error) {
//	db, err := sql.Open(driver, dsn)
//	if err != nil {
//		return nil, err
//	}
//
//	return &DB{
//		db:  db,
//		dsn: dsn,
//		log: log,
//	}, nil
//}
//
//func (d DB) Exec(query string, args ...interface{}) (sql.Result, err) {
//	d.log.Println(dsn, query, args)
//	return d.db.Exec(query, args...)
//}
