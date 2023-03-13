package numlookupapi

import (
	"errors"
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
	if len(number) < 12 {
		return errors.New("validation error: Кол-во цифр не должно быть меньше 12")
	} else if string(number[0]) != "+" {
		return errors.New("validation error: Номер должен начинаться со знака '+'")
	} else {
		return nil
	}
}
