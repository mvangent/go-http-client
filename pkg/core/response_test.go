package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpofe/go-http-client/pkg/gomime"
)

type Train struct {
	Name string `json:"name"`
	Cold bool   `json:"cold"`
}

func TestResponse(t *testing.T) {
	responseBody := Train{Name: "ICE", Cold: true}

	bytes, err := json.Marshal(responseBody)

	assert.Nil(t, err)

	headers := make(http.Header)
	headers.Set(gomime.HeaderUserAgent, "DB")

	response := Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       bytes,
		Headers:    headers,
	}

	var train Train

	assert.Nil(t, response.UnmarshalJson(&train))
	assert.Equal(t, responseBody, train)

	assert.Equal(t, `{"name":"ICE","cold":true}`, response.String())
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "200 OK", response.Status)
	assert.Equal(t, headers, response.Headers)
}
