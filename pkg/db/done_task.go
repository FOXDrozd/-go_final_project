package db

import (
	"database/sql"
)

func DoneTask(id string, dateStr *string, repeat *string) error {
	err := DB.QueryRow(
		"SELECT date, repeat FROM scheduler WHERE id = ?",
		id,
	).Scan(&dateStr, &repeat)

	if err == sql.ErrNoRows {
		return ErrTaskNotFound
	}

	if err != nil {
		return err
	}

	if *repeat == "" {
		_, err = DB.Exec("DELETE FROM scheduler WHERE id = ?", id)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}
