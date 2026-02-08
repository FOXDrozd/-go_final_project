package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeat := r.FormValue("repeat")

	// Нормализация строки repeat: удаляем лишние пробелы и заменяем плюсы на пробелы
	repeat = strings.ReplaceAll(repeat, "+", " ")
	repeat = strings.TrimSpace(repeat)
	repeat = strings.Join(strings.Fields(repeat), " ")

	now, err := time.Parse(layout, nowStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrCannotNextDate.Error()})
		return
	}

	dstart, err := time.Parse(layout, dateStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": ErrCannotNextDate.Error()})
		return
	}

	next, err := NextDate(now, dstart, repeat)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	fmt.Fprint(w, next)
}
