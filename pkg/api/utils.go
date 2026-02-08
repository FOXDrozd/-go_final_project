package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func parseIntList(s string, min, max int) []int {
	var res []int
	for _, p := range strings.Split(s, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil && n >= min && n <= max {
			res = append(res, n)
		}
	}
	return res
}

func parseMonthDays(s string) []int {
	var res []int
	for _, p := range strings.Split(s, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil && (n >= 1 && n <= 31 || n == -1 || n == -2) {
			res = append(res, n)
		}
	}
	return res
}

func contains(arr []int, v int) bool {
	for _, n := range arr {
		if n == v {
			return true
		}
	}
	return false
}

func lastDayOfMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
