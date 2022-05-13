package gohttp

import (
	"net/http"
	"time"
)

type ClientBuilder interface {
	Build() Client

	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder
	SetResponseTimeout(responseTimeout time.Duration) ClientBuilder
	SetMaxIdleConnections(maxIdleConnections int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
}

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (b *clientBuilder) Build() Client {
	client := httpClient{builder: b}
	return &client
}

func (b *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	b.headers = headers
	return b
}

func (b *clientBuilder) SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder {
	b.connectionTimeout = connectionTimeout
	return b
}

func (b *clientBuilder) SetResponseTimeout(responseTimeout time.Duration) ClientBuilder {
	b.responseTimeout = responseTimeout
	return b
}

func (b *clientBuilder) SetMaxIdleConnections(maxIdleConnections int) ClientBuilder {
	b.maxIdleConnections = maxIdleConnections
	return b
}

func (b *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	b.disableTimeouts = disable
	return b
}
