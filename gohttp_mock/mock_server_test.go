package gohttp_mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TravelInquiry struct {
	User        string `json:"user"`
	Destination string `json:"destination"`
	Departure   string `json:"departure"`
	Datetime    string `json:"datetime"`
}

var travelInquiry = TravelInquiry{
	User:        "vpofe",
	Destination: "Amsterdam",
	Departure:   "Berlin",
	Datetime:    "22052022-07H00M",
}

func TestSuccesfulMockedScenario(t *testing.T) {
	assert.NotNil(t, MockupServer)

	MockupServer.Start()

	assert.Equal(t, true, MockupServer.IsEnabled())

	var (
		method       = "POST"
		url          = "https://bahn.de/search"
		requestBody  = `{"user":"vpofe","destination":"Amsterdam", "departure": "Berlin", "datetime":"22052022-07H00M"}`
		statusCode   = 201
		responseBody = `{results: [IC2201, [ICE78, IC34]]}`
	)

	testMock := Mock{
		Method:      method,
		Url:         url,
		RequestBody: requestBody,

		ResponseStatusCode: statusCode,
		ResponseBody:       responseBody,
	}

	MockupServer.AddMock(testMock)

	client := MockupServer.GetClient()

	assert.NotNil(t, client)

	buf, err := json.Marshal(travelInquiry)

	assert.Nil(t, err)

	request, err := http.NewRequest(method, url, bytes.NewBuffer(buf))

	assert.Nil(t, err)
	assert.NotNil(t, request)

	response, err := client.Do(request)

	assert.Nil(t, err)

	assert.Equal(t, statusCode, response.StatusCode)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, responseBody, string(body))

	assert.Equal(t, 1, len(MockupServer.mocks))

	MockupServer.Flush()

	assert.Equal(t, 0, len(MockupServer.mocks))

	MockupServer.Stop()

	assert.Equal(t, false, MockupServer.IsEnabled())
}

func TestErrorMockedScenario(t *testing.T) {
	assert.NotNil(t, MockupServer)

	MockupServer.Start()

	assert.Equal(t, true, MockupServer.IsEnabled())

	var (
		method = "POST"
		url    = "https://bahn.de/search"
	)

	client := MockupServer.GetClient()

	assert.NotNil(t, client)

	buf, err := json.Marshal(travelInquiry)

	assert.Nil(t, err)

	request, err := http.NewRequest(method, url, bytes.NewBuffer(buf))

	assert.Nil(t, err)
	assert.NotNil(t, request)

	response, err := client.Do(request)

	assert.Nil(t, response)

	assert.NotNil(t, err)

	errorMsg := fmt.Sprintf("There is no mock matching %s from %s with given body \n", method, url)

	assert.Equal(t, err.Error(), errorMsg)

	assert.Equal(t, 0, len(MockupServer.mocks))
}
