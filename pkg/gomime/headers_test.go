package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Error("invalid Content-Type header")
	}

	if HeaderUserAgent != "User-Agent" {
		t.Error("invalid User-Agent header")
	}
}

func TestContentTypeValues(t *testing.T) {
	if ContentTypeJson != "application/json" {
		t.Error("invalid Content-Type value for application/json")
	}

	if ContentTypeXml != "application/xml" {
		t.Error("invalid Content-Type value for application/xml")
	}

	if ContentTypeOctetStream != "application/octet-stream" {
		t.Error("invalid Content-Type value for application/octet-stream")
	}
}
