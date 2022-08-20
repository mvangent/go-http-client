package gohttp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpofe/go-http-client/gomime"
)

func TestAllHeaders(t *testing.T) {
	var (
		headers         = make(http.Header)
		requestHeader   = make(http.Header)
		expectedHeaders = make(http.Header)
	)

	headers.Set("X-test", "test-value")
	requestHeader.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	expectedHeaders.Set("X-test", "test-value")
	expectedHeaders.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	expectedHeaders.Set(gomime.HeaderUserAgent, "local-machine")

	client := httpClient{
		builder: &clientBuilder{
			headers:   headers,
			userAgent: "local-machine",
		},
	}

	assert.Equal(t, expectedHeaders, client.getRequestHeaders(requestHeader))
}

func TestOnlyRequestHeader(t *testing.T) {
	var (
		requestHeader   = make(http.Header)
		expectedHeaders = make(http.Header)
	)

	requestHeader.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	expectedHeaders.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := httpClient{
		builder: &clientBuilder{
			headers: nil,
		},
	}

	assert.Equal(t, expectedHeaders, client.getRequestHeaders(requestHeader))
}

func TestOnlyStandardHeader(t *testing.T) {
	var (
		headers         = make(http.Header)
		expectedHeaders = make(http.Header)
	)

	headers.Set(gomime.HeaderUserAgent, "local-machine")

	expectedHeaders.Set(gomime.HeaderUserAgent, "local-machine")

	client := httpClient{
		builder: &clientBuilder{
			headers: headers,
		},
	}

	assert.Equal(t, expectedHeaders, client.getRequestHeaders(nil))
}

func TestGetHeadersHelper(t *testing.T) {
	var (
		firstHeaders  = make(http.Header)
		secondHeaders = make(http.Header)
	)

	firstHeaders.Set(gomime.HeaderContentType, gomime.ContentTypeXml)
	secondHeaders.Set(gomime.HeaderUserAgent, "Blechley Park")

	assert.Equal(t, firstHeaders, getHeaders(firstHeaders, secondHeaders))
	assert.Equal(t, make(http.Header), getHeaders())
}
