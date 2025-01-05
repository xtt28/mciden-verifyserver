package main

import (
	"log"

	"github.com/xtt28/mciden-verifyserver/config"
	"github.com/xtt28/mciden-verifyserver/db"
	"github.com/xtt28/mciden-verifyserver/server"
)

func main() {
	conf, err := config.ReadConfigFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	sqldb, err := db.Connect(conf.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(sqldb, conf.ServerURL)
	if err != nil {
		log.Fatal(err)
	}
}