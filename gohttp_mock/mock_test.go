package gohttp_mock

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponse(t *testing.T) {
	var (
		method      = "POST"
		url         = "https://bahn.de/search"
		requestBody = `Invalid Body`
	)

	testErrorMock := Mock{
		Method:      method,
		Url:         url,
		RequestBody: requestBody,

		Error:              errors.New("Mocked error"),
		ResponseStatusCode: 500,
		ResponseBody:       "error error oh no",
	}

	response, err := testErrorMock.GetResponse()

	assert.Equal(t, "Mocked error", err.Error())
	assert.Nil(t, response)
}

func TestSuccessResponse(t *testing.T) {
	var (
		method       = "POST"
		url          = "https://bahn.de/search"
		requestBody  = `{"user":"vpofe","destination":"Amsterdam", "departure": "Berlin", "datetime":"22052022-07H00M"}`
		statusCode   = 201
		responseBody = `{results: [IC2201, [ICE78, IC34]]}`
	)

	testSuccessMock := Mock{
		Method:      method,
		Url:         url,
		RequestBody: requestBody,

		ResponseStatusCode: statusCode,
		ResponseBody:       responseBody,
	}

	response, err := testSuccessMock.GetResponse()

	assert.Nil(t, err)

	assert.Equal(t, responseBody, response.String())
	assert.Equal(t, statusCode, response.StatusCode)
}
