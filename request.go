package numlookupapi

import (
	"context"
	"errors"
	"net/http"
	"nlkapi/pkg/check"
	"time"
)

type Params struct {
	ApiKey string
}

// GetRequest запрос данных
func (m *Params) GetRequest(nubmer string) (Response, error) {
	if err := check.PhoneNumberStruct(nubmer); err != nil {
		return Response{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.numlookupapi.com/v1/validate/"+nubmer, nil)
	if err != nil {
		return Response{}, err
	}

	request.Header.Add("apikey", m.ApiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Response{}, err
	}

	switch response.StatusCode {
	case 200:
		return Response{response}, nil
	case 401:
		return Response{}, errors.New("invalid authentication credentials")
	case 403:
		return Response{}, errors.New("you are not allowed to use this endpoint, please upgrade your plan")
	case 404:
		return Response{}, errors.New("a requested endpoint does not exist")
	case 422:
		return Response{}, errors.New("validation error")
	case 429:
		return Response{}, errors.New("you have hit your rate limit or your monthly limit")
	case 500:
		return Response{}, errors.New("internal Server Error - let us know: support@numlookupapi.com")
	default:
		return Response{}, errors.New("error Process")
	}
}
