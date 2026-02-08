package api

import (
	"net/http"
	"strconv"
	"time"

	"main.go/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	limit := 50
	search := r.URL.Query().Get("search")

	tasks, err := Tasks(limit, search)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// На всякий случай, если tasks nil
	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJSON(w, http.StatusOK, TasksResp{Tasks: tasks})
}

func Tasks(limit int, search string) ([]*db.Task, error) {
	var query string
	var args []any

	if search != "" {
		if t, err := time.Parse(layout, search); err == nil {

			dateStr := t.Format(layout)
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

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, ErrCannotReadRow
	}

	defer rows.Close()

	tasks := make([]*db.Task, 0)
	for rows.Next() {
		var t db.Task
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
