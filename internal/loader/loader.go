package loader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type loader struct {
}

type Loader interface {
	LoadPage(url string) ([]byte, error)
}

func NewLoader() Loader {
	return &loader{}
}

func (l * loader) LoadPage(url string) ([]byte, error) {
	response, err := HandleResponse(url)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	return bytes, nil
}

func HandleResponse(url string) (*http.Response, error) {
	for {
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		switch {
		case isSuccess(response.StatusCode):
			return response, nil

		case tooManyRequests(response.StatusCode):
			time.Sleep(15 * time.Second)
			continue

		default:
			return nil, fmt.Errorf("Response status %d ", response.StatusCode)
		}
	}
}

func isSuccess(code int) bool {
	return code == 200
}

func tooManyRequests(code int) bool {
	return code == 429
}
