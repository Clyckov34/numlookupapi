package numlookupapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Body struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

type Response struct {
	*http.Response
}

// Data Выгрузка данных в структуру
func (m *Response) Data() (Body, error) {
	body, err := io.ReadAll(m.Body)
	if err != nil {
		return Body{}, err
	}

	var bodyJson Body
	if err := json.Unmarshal(body, &bodyJson); err != nil {
		return Body{}, err
	}

	return bodyJson, nil
}

// PhoneValid валидность номера.
// Образец: true
func (m *Body) PhoneValid() bool {
	return m.Valid
}

// PhoneNumber номер телефона
// Образец: "79998887766"
func (m *Body) PhoneNumber() string {
	return m.Number
}

// PhoneLocalFormat() номер телефона без кода страны
// Образец: "9998887766"
func (m *Body) PhoneLocalFormat() string {
	return m.LocalFormat
}

// PhoneInternationalFormat международный формат
// Образец: "+79998887766"
func (m *Body) PhoneInternationalFormat() string {
	return m.InternationalFormat
}

// PhoneCountryPrefix префикс код страны
// Образец: "+7"
func (m *Body) PhoneCountryPrefix() string {
	return m.CountryPrefix
}

// PhoneCountryCode код страны
// Образец: "RU"
func (m *Body) PhoneCountryCode() string {
	return m.CountryCode
}

// PhoneCountryName название страны
// Образец: "Russian Federation"
func (m *Body) PhoneCountryName() string {
	return m.CountryName
}

// PhoneLocation расположения зарегистрированного номера
// Образец: "Volgograd Oblast"
func (m *Body) PhoneLocation() string {
	return m.Location
}

// PhoneСarrier сотовый оператор
// Образец: "LLC Skartel (YOTA)"
func (m *Body) PhoneСarrier() string {
	return m.Carrier
}

// PhoneLineType тип связи
// Образец: "mobile"
func (m *Body) PhoneLineType() string {
	return m.LineType
}
