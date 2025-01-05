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
			return c.Render(http.StatusNotFound, "error.html", "Invalid verification intent ID.")
		}
		if err != nil {
			c.Logger().Error(err)
			return c.Render(http.StatusInternalServerError, "error.html", "An error occurred while finding your profile.")
		}

		if intent.IsExpired() {
			return c.Render(http.StatusGone, "error.html", "This verification link has expired.")
		}

		err = db.VerifyAccount(sqldb, intent)
		if err != nil {
			c.Logger().Error(err)
			if me, ok := err.(*mysql.MySQLError); ok {
				if me.Number == db.ErrDuplicate {
					return c.Render(http.StatusConflict, "error.html", "You already have an account registered with this email address.")
				}
			}

			return c.Render(http.StatusInternalServerError, "error.html", "An error occurred while making your profile.")
		}

		student, err := db.GetStudentByID(sqldb, intent.StudentID)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.html", "Couldn't look up student information. Your account still should have been activated.")
		}
		c.Logger().Info("Verified Minecraft account (UUID: %s) with intent %s belonging to %s %s", intent.PlayerUUID, intent.ID, student.FirstName, student.LastName)
		return c.Render(http.StatusOK, "confirm.html", student)
	}
}
