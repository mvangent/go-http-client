package gohttp

import (
	"net/http"
	"testing"

	"github.com/vpofe/go-http-client/gomime"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{builder: &clientBuilder{}}

	commonHeaders := make(http.Header)
	commonHeaders.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	commonHeaders.Set(gomime.HeaderUserAgent, "vpofes-http-client")

	client.builder.headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "123-456")

	// Execution
	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("We expect 3 headers")
	}

	if finalHeaders.Get(gomime.HeaderContentType) != gomime.ContentTypeJson {
		t.Error("Invalid Content-Type received")
	}

	if finalHeaders.Get(gomime.HeaderUserAgent) != "vpofes-http-client" {
		t.Error("Invalid User-Agent received")
	}

	if finalHeaders.Get("X-Request-Id") != "123-456" {
		t.Error("Invalid X-Request-Id received")
	}
}

func TestGetRequestBodyNilBody(t *testing.T) {
	// Initialization
	client := httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		requestBody, err := client.getRequestBody(gomime.ContentTypeJson, nil)

		if err != nil {
			t.Error("No error expected when passing a nil body")
		}

		if requestBody != nil {
			t.Error("No body expected when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody(gomime.ContentTypeJson, requestBody)

		if err != nil {
			t.Error("No error expected when marshalling a slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("Invalid json body when marshalling slice as json")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody(gomime.ContentTypeXml, requestBody)

		if err != nil {
			t.Error("No error expected when marshalling a slice as xml")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			t.Error("Invalid xml body when marshalling slice as xml")
		}
	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("", requestBody)

		if err != nil {
			t.Error("No error expected when marshalling a slice as json by default")
		}

		if string(body) != `["one","two"]` {
			t.Error("Invalid json body when marshalling slice as json by default")
		}
	})

}
