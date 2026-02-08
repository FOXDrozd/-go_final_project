package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"main.go/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrInvalidJSON.Error()})
		return
	}

	// title обязателен
	if task.Title == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrEmptyTitle.Error()})
		return
	}

	// Проверка и корректировка даты
	if err := checkDate(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Добавляем задачу в базу
	id, err := addTask(&task)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"id": string(id)})
}

func addTask(task *db.Task) (string, error) {
	query := `
		INSERT INTO scheduler (date, title, comment, repeat)
		VALUES (?, ?, ?, ?)
	`
	res, err := db.DB.Exec(
		query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
	)
	if err != nil {
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", ErrCannotGetLastID
	}

	return strconv.FormatInt(id, 10), nil
}
