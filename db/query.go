package db

import "database/sql"

func GetStudentByID(db *sql.DB, id int64) (res Student, err error) {
	rows, err := db.Query("SELECT * FROM `students` WHERE `id` = ?", id)
	if err != nil {
		return
	}

	defer rows.Close()
	if !rows.Next() {
		return
	}

	if err = rows.Scan(&res.ID, &res.FirstName, &res.LastName, &res.Email, &res.Academy); err != nil {
		return
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}