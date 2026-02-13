package api

import (
	"net/http"
	"time"

	"main.go/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	limit := limit
	search := r.URL.Query().Get("search")

	var dateStr string
	t, err := time.Parse(layout, search)
	if err != nil {
		dateStr = t.Format(layout)
	}

	tasks, err := db.GetTasks(limit, search, dateStr)
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
