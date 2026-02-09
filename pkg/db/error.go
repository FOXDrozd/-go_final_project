package db

import "errors"

var (
	ErrCannotInitDatabase = errors.New("Невозможно инициализировать базу данных")
)
