package db

import "database/sql"

func VerifyAccount(db *sql.DB, intent VerifyIntent) (err error) {
	_, err = db.Exec("INSERT INTO `registrations` (`student_id`, `player_uuid`) VALUES (?, ?)", intent.StudentID, intent.PlayerUUID)
	return
}