package api

import "net/http"

func Init() {
	http.HandleFunc("/api/signin", signInHandler)
	http.HandleFunc("/api/nextdate", auth(nextDayHandler))
	http.HandleFunc("/api/task", auth(getTaskHandler))
	http.HandleFunc("/api/tasks", auth(getTasksHandler))
	http.HandleFunc("/api/task/done", auth(taskDoneHandler))
}
