package api

import (
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

	err := db.DoneTask(id, &dateStr, &repeat)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
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

	err = db.UpdateDateTask(nextDate, id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}
