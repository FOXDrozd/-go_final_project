package db

import (
	"database/sql"
	"strconv"
)

func GetTask(id string) (*Task, error) {

	query := `
		SELECT id, date, title, comment, repeat
		FROM scheduler
		WHERE id = ?
	`

	task := &Task{}
	var dbID int64
	err := DB.QueryRow(query, id).Scan(
		&dbID,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	task.ID = strconv.FormatInt(dbID, 10)
	return task, nil
}
