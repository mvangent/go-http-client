package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)

	fullHeaders := c.getRequestHeaders(headers)

	request.Header = fullHeaders

	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	return client.Do(request)
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers to the request
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Add custom headers to the request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}
