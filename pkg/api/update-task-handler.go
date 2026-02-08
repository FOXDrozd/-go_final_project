package api

import (
	"encoding/json"
	"net/http"

	"main.go/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrInvalidJSON.Error()})
		return
	}

	// id обязателен
	if task.ID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrEmptyID.Error()})
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

	// Обновление задачи в базу
	if err := updateTask(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// успешный ответ — пустой JSON
	writeJSON(w, http.StatusOK, map[string]string{})
}

func updateTask(task *db.Task) error {
	if task.ID == "" {
		return ErrEmptyID
	}

	query := `UPDATE scheduler SET date=?, title=?, comment=?, repeat=? WHERE id=?`
	res, err := db.DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrTaskNotFound
	}

	return nil
}
