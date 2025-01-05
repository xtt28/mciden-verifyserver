package config

import "github.com/go-sql-driver/mysql"

type Config struct {
	DBConfig mysql.Config `json:"dbConfig"`
	ServerURL string `json:"serverURL"`
}