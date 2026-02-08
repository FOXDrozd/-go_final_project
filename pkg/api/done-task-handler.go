package api

import (
	"database/sql"
	"net/http"
	"time"

	"main.go/pkg/db"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrRequiredField.Error()})
		return
	}

	var dateStr, repeat string

	err := db.DB.QueryRow(
		"SELECT date, repeat FROM scheduler WHERE id = ?",
		id,
	).Scan(&dateStr, &repeat)

	if err == sql.ErrNoRows {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": ErrTaskNotFound.Error(),
		})
		return
	}

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if repeat == "" {
		_, err = db.DB.Exec("DELETE FROM scheduler WHERE id = ?", id)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{})
		return
	}

	date, err := time.Parse(layout, dateStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrInvalidDate.Error()})
		return
	}

	nextDate, err := NextDate(time.Now(), date, repeat)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	_, err = db.DB.Exec(
		"UPDATE scheduler SET date = ? WHERE id = ?",
		nextDate,
		id,
	)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}
