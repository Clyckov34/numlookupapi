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

// parserJSON Выгрузка данных в структуру
func parserJSON(r *http.Response) (*Body, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var bodyJson Body
	if err := json.Unmarshal(body, &bodyJson); err != nil {
		return nil, err
	}

	return &bodyJson, nil
}
