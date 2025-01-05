package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/xtt28/mciden-verifyserver/db"
)

func Activate(sqldb *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		intentID := c.Param("id")

		intent, err := db.GetVerifyIntentByID(sqldb, intentID)
		if intent == (db.VerifyIntent{}) {
			return c.String(http.StatusNotFound, "Invalid verification intent ID.")
		}
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, "An error occurred while finding your profile.")
		}

		if intent.IsExpired() {
			return c.String(http.StatusGone, "This verification link has expired.")
		}

		err = db.VerifyAccount(sqldb, intent)
		if err != nil {
			c.Logger().Error(err)
			if me, ok := err.(*mysql.MySQLError); ok {
				if me.Number == db.ErrDuplicate {
					return c.String(http.StatusConflict, "You already have an account registered with this email address.")
				} else {
					return c.String(http.StatusInternalServerError, "An error occurred while making your profile.")
				}
			} else {
				return c.String(http.StatusInternalServerError, "An error occurred while making your profile.")
			}
		}

		return c.String(http.StatusOK, "Done! Log into Minecraft to access the server.")
	}
}
