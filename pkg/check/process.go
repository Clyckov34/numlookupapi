package check

import (
	"errors"
	"strings"
)

// PhoneNumberStruct проверка формат номера телефона
func PhoneNumberStruct(number string) error {
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
