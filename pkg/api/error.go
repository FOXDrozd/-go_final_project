package api

import "errors"

var (
	//Ошибкт повтора
	ErrInvalidRepeat  = errors.New("Неверное значение repeat")
	ErrNoRepeat       = errors.New("Нет значения repeat")
	ErrCannotNextDate = errors.New("Невозможно вычислить следующую дату")

	//Ошибки форматов данных
	ErrInvalidDate = errors.New("Неверный формат даты")
	ErrInvalidJSON = errors.New("Неверный формат JSON")

	//Ошибка метода
	ErrMethodNotAllowed = errors.New("Метод не поддерживается")

	ErrEmptyTitle = errors.New("Не указан заголовок задачи")
	ErrEmptyID    = errors.New("Не указан идентификатор")

	ErrRequiredField = errors.New("Отсутствует обязательное поле")

	ErrPassword        = errors.New("Неверный пароль")
	ErrInvalidToken    = errors.New("Неверный токен")
	ErrTokenGeneration = errors.New("Ошибка генерации токена")
)
