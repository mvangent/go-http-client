package gohttp_mock

import (
	"fmt"
	"net/http"

	"github.com/vpofe/go-http-client/pkg/core"
)

type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseBody       string
	ResponseStatusCode int
}

func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := core.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Headers:    map[string][]string{},
		Body:       []byte(m.ResponseBody),
	}

	return &response, nil
}
