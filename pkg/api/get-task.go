package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"main.go/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": ErrRequiredField.Error(),
		})
		return
	}

	task, err := getTask(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func getTask(id string) (*db.Task, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	query := `
		SELECT id, date, title, comment, repeat
		FROM scheduler
		WHERE id = ?
	`

	task := &db.Task{}
	var dbID int64
	err := db.DB.QueryRow(query, id).Scan(
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
