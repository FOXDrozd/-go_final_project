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

	if task.ID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrEmptyID.Error()})
		return

	}
	// Обновление задачи в базу
	if err := db.UpdateTask(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// успешный ответ — пустой JSON
	writeJSON(w, http.StatusOK, map[string]string{})
}
