package gohttp

import (
	"net/http"

	"github.com/vpofe/go-http-client/gomime"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}

	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers to the request
	for header, value := range c.builder.headers {
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

	// Adding User-Agent from the client builder if custom or common headers have not defined it yet
	if c.builder.userAgent != "" {
		if result.Get("User-Agent") != "" {
			return result
		}

		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}

	return result
}
