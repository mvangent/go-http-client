package gohttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	singletonClient = NewBuilder().Build()
)

type TestRequestBody struct {
	Id       int     `json:"Id"`
	Customer string  `json:"Customer"`
	Quantity int     `json:"Quantity"`
	Price    float32 `json:"Price"`
}

func TestClientGet(t *testing.T) {
	response, err := singletonClient.Get("https://reqbin.com/echo")

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, 200, response.StatusCode)
}

func TestClientPut(t *testing.T) {
	requestBody := TestRequestBody{
		12345,
		"John Smith",
		1,
		10.00}

	response, err := singletonClient.Put("https://reqbin.com/echo/put/json", requestBody)

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, 200, response.StatusCode)

	type Status struct {
		Success string `json:"success"`
	}

	var status Status

	unmarshalErr := response.UnmarshalJson(&status)

	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "true", status.Success)
}

func TestClientPost(t *testing.T) {

	requestBody := TestRequestBody{
		78912,
		"Jason Sweet",
		1,
		18.00}

	response, err := singletonClient.Post("https://reqbin.com/echo/post/json", requestBody)

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, 200, response.StatusCode)

	type Status struct {
		Success string `json:"success"`
	}

	var status Status

	unmarshalErr := response.UnmarshalJson(&status)

	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "true", status.Success)
}

func TestClientPatch(t *testing.T) {
	requestBody := TestRequestBody{
		12345,
		"John Smith",
		1,
		10.00}

	response, err := singletonClient.Patch("https://reqbin.com/echo/patch/json", requestBody)

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, 200, response.StatusCode)

	type Status struct {
		Success string `json:"success"`
	}

	var status Status

	unmarshalErr := response.UnmarshalJson(&status)

	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "true", status.Success)
}

func TestClientDelete(t *testing.T) {
	response, err := singletonClient.Delete("https://reqbin.com/sample/delete/json")

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, 200, response.StatusCode)

	type Status struct {
		Success string `json:"success"`
	}

	var status Status

	unmarshalErr := response.UnmarshalJson(&status)

	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "true", status.Success)
}
