package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/xtt28/mciden-verifyserver/db"
)

func Activate(sqldb *sql.DB) echo.HandlerFunc {
	logError := func(c echo.Context, err error) {
		c.Logger().Errorj(log.JSON{
			"context": "Could not activate user account.",
			"err": err,
		})
	}

	renderError := func(c echo.Context, code int, description string) error {
		return c.Render(code, "error.html", description)
	}

	return func(c echo.Context) error {
		intentID := c.Param("id")

		intent, err := db.GetVerifyIntentByID(sqldb, intentID)
		
		if err != nil {
			logError(c, err)
			return renderError(c, http.StatusNotFound, "An error occurred while finding your profile.")
		}
		
		if intent.IsEmpty() {
			return renderError(c, http.StatusNotFound, "Invalid verification intent ID.")
		}

		if intent.IsExpired() {
			return renderError(c, http.StatusGone, "This verification link has expired.")
		}

		err = db.VerifyAccount(sqldb, intent)
		if err != nil {
			logError(c, err)
			if me, ok := err.(*mysql.MySQLError); ok {
				if me.Number == db.ErrDuplicate {
					return renderError(c, http.StatusConflict, "You already have an account registered with this email address.")
				}
			}

			return renderError(c, http.StatusInternalServerError, "An error occurred while making your profile.")
		}

		student, err := db.GetStudentByID(sqldb, intent.StudentID)
		if err != nil {
			logError(c, err)
			return renderError(c, http.StatusInternalServerError, "Couldn't look up student infoormation. Your account still should have been activated.")
		}
		c.Logger().Printf("Verified Minecraft account (UUID: %s) with intent %s belonging to %s %s", intent.PlayerUUID, intent.ID, student.FirstName, student.LastName)
		return c.Render(http.StatusOK, "confirm.html", student)
	}
}
