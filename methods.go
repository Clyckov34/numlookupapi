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