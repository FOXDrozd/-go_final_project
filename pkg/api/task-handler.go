package api

import (
	"net/http"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTaskHandler(w, r)

	case http.MethodPost:
		addTaskHandler(w, r)

	case http.MethodPut:
		updateTaskHandler(w, r)

	case http.MethodDelete:
		taskDeleteHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": ErrMethodNotAllowed.Error()})
	}
}
