package db

import (
	"strconv"
)

func GetTasks(limit int, search string, dateStr string) ([]*Task, error) {
	var query string
	var args []any

	if search != "" {

		if dateStr != "" {
			query = `
				SELECT id, date, title, comment, repeat
				FROM scheduler
				WHERE date = ?
				ORDER BY date ASC
				LIMIT ?
			`
			args = []any{dateStr, limit}
		} else {
			query = `
				SELECT id, date, title, comment, repeat
				FROM scheduler
				WHERE title LIKE ? OR comment LIKE ?
				ORDER BY date ASC
				LIMIT ?
			`
			like := "%" + search + "%"
			args = []any{like, like, limit}
		}
	} else {
		query = `
			SELECT id, date, title, comment, repeat
			FROM scheduler
			ORDER BY date ASC
			LIMIT ?
		`
		args = []any{limit}
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, ErrCannotReadRow
	}

	defer rows.Close()

	tasks := make([]*Task, 0)
	for rows.Next() {
		var t Task
		var id int64
		if err := rows.Scan(&id, &t.Date, &t.Title, &t.Comment, &t.Repeat); err != nil {
			return nil, ErrCannotReadRow
		}
		t.ID = strconv.FormatInt(id, 10)
		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, ErrCannotReadRow
	}

	return tasks, nil
}
