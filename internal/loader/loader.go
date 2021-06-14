package loader

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	err = HandleResponseStatus(response.StatusCode)
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

func HandleResponseStatus(statusCode int) error {
	switch {
	case isSuccess(statusCode):
		return nil

	case isClientError(statusCode):
		return fmt.Errorf("Response status %d ", statusCode)

	case isServerError(statusCode):
		return fmt.Errorf("Response status %d ", statusCode)

	default:
		return nil
	}
}

func isSuccess(code int) bool {
	return code == 200
}

func isClientError(code int) bool {
	return code >= 400 && code < 500
}

func isServerError(code int) bool {
	return code >= 500
}
