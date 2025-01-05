package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/xtt28/mciden-verifyserver/handler"
)

func Start(db *sql.DB, listenAddr string) error {
	e := echo.New()
	e.GET("/activate/:id", handler.Activate(db))
	return e.Start(listenAddr)
}