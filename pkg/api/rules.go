package api

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, dstart time.Time, repeat string) (string, error) {
	if repeat == "" {
		return "", ErrNoRepeat
	}

	parts := strings.Fields(repeat)

	switch parts[0] {
	case "d":
		return NextByDays(now, dstart, parts)
	case "w":
		return NextByWeekdays(now, parts)
	case "m":
		return NextByMonths(now, parts)
	case "y":
		return NextByYear(now, dstart)
	}

	return "", ErrInvalidRepeat
}

func NextByDays(now time.Time, start time.Time, parts []string) (string, error) {
	if len(parts) != 2 {
		return "", ErrInvalidRepeat
	}

	days, err := strconv.Atoi(parts[1])
	if err != nil || days < 1 || days > 400 {
		return "", ErrInvalidRepeat
	}

	next := start.AddDate(0, 0, days)

	for !next.After(now) {
		next = next.AddDate(0, 0, days)
	}

	return next.Format(layout), nil
}

func NextByWeekdays(now time.Time, parts []string) (string, error) {

	if len(parts) != 2 {
		return "", ErrInvalidRepeat
	}

	days := parseIntList(parts[1], 1, 7)
	if len(days) == 0 {
		return "", ErrInvalidRepeat
	}

	sort.Ints(days)

	for i := 1; i <= 14; i++ {
		dateWD := now.AddDate(0, 0, i)
		wd := int(dateWD.Weekday())
		if wd == 0 {
			wd = 7
		}
		if contains(days, wd) {
			return dateWD.Format(layout), nil
		}
	}

	return "", ErrCannotNextDate
}

func NextByMonths(now time.Time, parts []string) (string, error) {
	if len(parts) < 2 || len(parts) > 3 {
		return "", ErrInvalidRepeat
	}

	mdays := parseMonthDays(parts[1])
	if len(mdays) == 0 {
		return "", ErrInvalidRepeat
	}

	var months []int
	if len(parts) == 3 {
		months = parseIntList(parts[2], 1, 12)
		if len(months) == 0 {
			return "", ErrInvalidRepeat
		}
	}

	check := now.AddDate(0, 0, 1)

	for i := 0; i < 370; i++ {
		y, m, _ := check.Date()
		if len(months) > 0 && !contains(months, int(m)) {
			check = check.AddDate(0, 0, 1)
			continue
		}

		lastDay := lastDayOfMonth(y, m)

		for _, d := range mdays {
			var day int
			if d > 0 {
				day = d
			} else {
				day = lastDay + d + 1
			}

			if day == check.Day() {
				return check.Format(layout), nil
			}
		}

		check = check.AddDate(0, 0, 1)
	}

	return "", ErrCannotNextDate
}

func NextByYear(now time.Time, start time.Time) (string, error) {
	next := start.AddDate(1, 0, 0) // сразу следующий год

	// Если после добавления года всё ещё <= now, прибавляем еще годы
	for !next.After(now) {
		next = next.AddDate(1, 0, 0)
	}

	return next.Format(layout), nil
}
