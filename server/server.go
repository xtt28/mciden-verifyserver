package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xtt28/mciden-verifyserver/handler"
	"github.com/xtt28/mciden-verifyserver/renderer"
)

func Start(db *sql.DB, listenAddr string) error {
	e := echo.New()
	e.HideBanner = true
	e.Renderer = renderer.NewRendererFromTarget("templates/*.html")
	e.Use(middleware.Logger())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.GET("/activate/:id", handler.Activate(db))

	return e.Start(listenAddr)
}