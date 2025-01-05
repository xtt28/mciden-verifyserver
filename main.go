package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/xtt28/mciden-verifyserver/db"
	"github.com/xtt28/mciden-verifyserver/server"
)

func main() {
	dbconf := mysql.Config{
		User: "root",
		Passwd: "",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "identityservice",
		AllowNativePasswords: true,
		ParseTime: true,
	}
	sqldb, err := db.Connect(dbconf)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(sqldb, ":8080")
	if err != nil {
		log.Fatal(err)
	}
}