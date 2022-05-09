package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "vpofes-http-client")

	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "123-456")

	// Execution
	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("We expect 3 headers")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid Content-Type received")
	}

	if finalHeaders.Get("User-Agent") != "vpofes-http-client" {
		t.Error("Invalid User-Agent received")
	}

	if finalHeaders.Get("X-Request-Id") != "123-456" {
		t.Error("Invalid X-Request-Id received")
	}
}
