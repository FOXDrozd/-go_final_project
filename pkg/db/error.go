package db

import "errors"

var (
	ErrCannotInitDatabase = errors.New("Невозможно инициализировать базу данных")
	ErrTaskNotFound       = errors.New("Задача не найдена")
	ErrRequiredField      = errors.New("Отсутствует обязательное поле")

	ErrCannotGetLastID = errors.New("Невозможно получить ID последней вставленной записи")
	ErrCannotReadRow   = errors.New("Невозможно прочитать строку из базы данных")
)
