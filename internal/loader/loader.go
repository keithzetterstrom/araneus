package loader

import (
	logtool "github.com/keithzetterstrom/araneus/tools/logger"
	"io/ioutil"
	"net/http"
)

type loader struct {
	Logger logtool.Logger
}

type Loader interface {
	LoadPage(url string) ([]byte, error)
}

func NewLoader(logger logtool.Logger) Loader {
	return &loader{
		Logger: logger,
	}
}

func (l * loader) LoadPage(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		l.Logger.ErrorLogger.Println(err)
		return nil, err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		l.Logger.ErrorLogger.Println(err)
		return nil, err
	}

	response.Body.Close()

	return bytes, nil
}
