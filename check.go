package numlookupapi

import (
	"errors"
	"strings"
)

// checkAPIKey проверка ключа
func checkAPIKey(apiKey string) error {
	if len(apiKey) == 0 {
		return errors.New("validation error: Нет в наличии ключа")
	}

	return nil
}

// checkPhoneNumber проверка формат номера телефона
func checkPhoneNumber(number string) error {
	split := strings.Split(number, "")

	switch {
	case split[0] != "+":
		return errors.New("validation error: Номер должен начинаться со знака '+'")
	case len(number) < 12:
		return errors.New("validation error: Кол-во цифр не должно быть меньше 12")
	default:
		return nil
	}
}
