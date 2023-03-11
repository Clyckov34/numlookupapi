package check

import (
	"errors"
	"strings"
)

// PhoneNumberStruct проверка формат номера телефона
func PhoneNumberStruct(number string) error {
	split := strings.Split(number, "")

	if split[0] != "+" {
		return errors.New("Validation error")
	} else if len(number) < 12 {
		return errors.New("Validation error")
	} else {
		return nil
	}
}