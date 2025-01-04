package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func Connect(conf mysql.Config) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", conf.FormatDSN())
	return
}