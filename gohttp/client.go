package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client *http.Client

	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	Headers            http.Header
}

func New() HttpClient {
	httpClient := &httpClient{}
	return httpClient
}

type HttpClient interface {
	SetHeaders(headers http.Header)
	SetConnectionTimeout(connectionTimeout time.Duration)
	SetResponseTimeout(responseTimeout time.Duration)
	SetMaxIdleConnections(maxIdleConnections int)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) SetConnectionTimeout(connectionTimeout time.Duration) {
	c.connectionTimeout = connectionTimeout
}

func (c *httpClient) SetResponseTimeout(responseTimeout time.Duration) {
	c.responseTimeout = responseTimeout
}

func (c *httpClient) SetMaxIdleConnections(maxIdleConnections int) {
	c.maxIdleConnections = maxIdleConnections
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)

}
func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)

}
func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)

}
func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
