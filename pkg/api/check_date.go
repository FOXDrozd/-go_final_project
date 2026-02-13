package api

import (
	"time"

	"main.go/pkg/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()
	today := now.Format(layout)

	// если дата пустая, ставим сегодня
	if task.Date == "" {
		task.Date = today
	}

	date, err := time.Parse(layout, task.Date)
	// проверяем формат даты
	if err != nil {
		return ErrInvalidDate
	}

	// если дата в прошлом, ставим сегодня
	if date.Before(now) {
		task.Date = today
	}

	// если repeat указан, игнорируем исходную дату и ставим сегодня
	if task.Repeat != "" {
		if _, err := NextDate(now, date, task.Repeat); err != nil {
			return err
		}

		task.Date = today
	}

	return nil
}
