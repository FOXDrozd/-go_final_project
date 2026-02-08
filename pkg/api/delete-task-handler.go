package api

import (
	"net/http"

	"main.go/pkg/db"
)

func taskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrRequiredField.Error()})
		return
	}

	res, err := db.DB.Exec("DELETE FROM scheduler WHERE id = ?", id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if rows == 0 {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": ErrTaskNotFound.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}
