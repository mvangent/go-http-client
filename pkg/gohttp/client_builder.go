package gohttp

import (
	"net/http"
	"time"
)

type ReadOnlyBuilderConfig struct {
	Headers            http.Header
	MaxIdleConnections int
	ConnectionTimeout  time.Duration
	ResponseTimeout    time.Duration
	DisableTimeouts    bool
	UserAgent          string
	CustomClient       http.Client
}

type ClientBuilder interface {
	Build() Client
	ReadBuilderConfig() ReadOnlyBuilderConfig

	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder
	SetResponseTimeout(responseTimeout time.Duration) ClientBuilder
	SetMaxIdleConnections(maxIdleConnections int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder
}

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
	client             *http.Client
	userAgent          string
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (cb *clientBuilder) Build() Client {
	client := httpClient{builder: cb}
	return &client
}

func (cb *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	cb.headers = headers
	return cb
}

func (cb *clientBuilder) SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder {
	cb.connectionTimeout = connectionTimeout
	return cb
}

func (b *clientBuilder) SetResponseTimeout(responseTimeout time.Duration) ClientBuilder {
	b.responseTimeout = responseTimeout
	return b
}

func (cb *clientBuilder) SetMaxIdleConnections(maxIdleConnections int) ClientBuilder {
	cb.maxIdleConnections = maxIdleConnections
	return cb
}

func (cb *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	cb.disableTimeouts = disable
	return cb
}

func (cb *clientBuilder) SetHttpClient(c *http.Client) ClientBuilder {
	cb.client = c

	return cb
}

func (cb *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	cb.userAgent = userAgent
	return cb
}

func (cb *clientBuilder) ReadBuilderConfig() ReadOnlyBuilderConfig {
	readOnlyConfig := ReadOnlyBuilderConfig{
		Headers:            cb.headers,
		MaxIdleConnections: cb.maxIdleConnections,
		ConnectionTimeout:  cb.connectionTimeout,
		ResponseTimeout:    cb.responseTimeout,
		DisableTimeouts:    cb.disableTimeouts,
		UserAgent:          cb.userAgent,
		CustomClient:       *cb.client,
	}

	return readOnlyConfig
}
