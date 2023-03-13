package numlookupapi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	ApiKey string
}

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

// GetResponse получить ответ от API
func (m *Client) GetResponse(nubmer string) (*Body, error) {
	if err := checkPhoneNumber(nubmer); err != nil {
		return nil, err
	}

	if err := checkAPIKey(m.ApiKey); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.numlookupapi.com/v1/validate/"+nubmer, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("apikey", m.ApiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case 200:
		res, err := parserJSON(response)
		if err != nil {
			return nil, err
		}

		return res, nil
	case 401:
		return nil, errors.New("invalid authentication credentials")
	case 403:
		return nil, errors.New("you are not allowed to use this endpoint, please upgrade your plan")
	case 404:
		return nil, errors.New("a requested endpoint does not exist")
	case 422:
		return nil, errors.New("validation error")
	case 429:
		return nil, errors.New("you have hit your rate limit or your monthly limit")
	case 500:
		return nil, errors.New("internal Server Error - let us know: support@numlookupapi.com")
	default:
		return nil, errors.New("error Process")
	}
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
