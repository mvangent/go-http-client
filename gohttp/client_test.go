package gohttp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMaxIdleConnections(t *testing.T) {
	var (
		maxIdleConnections = 8
	)

	client := httpClient{
		builder: &clientBuilder{
			maxIdleConnections: maxIdleConnections,
			disableTimeouts:    false,
		},
	}

	assert.Equal(t, maxIdleConnections, client.getMaxIdleConnections())

	client.builder.maxIdleConnections = 0

	assert.Equal(t, defaultMaxIdleConnections, client.getMaxIdleConnections())

}

func TestConnectionTimeout(t *testing.T) {
	var (
		connectionTimeout = 2 * time.Second
	)

	client := httpClient{
		builder: &clientBuilder{
			connectionTimeout: connectionTimeout,
		},
	}

	assert.Equal(t, connectionTimeout, client.getConnectionTimeout())

	client.builder.disableTimeouts = true

	assert.Equal(t, 0*time.Second, client.getConnectionTimeout())

	client.builder.connectionTimeout = 0 * time.Second
	client.builder.disableTimeouts = false

	assert.Equal(t, defaultConnectionTimeout, client.getConnectionTimeout())
}

func TestReponseTimeout(t *testing.T) {
	var (
		responseTimeout = 20 * time.Second
	)

	client := httpClient{
		builder: &clientBuilder{
			responseTimeout: responseTimeout,
		},
	}

	assert.Equal(t, responseTimeout, client.getResponseTimeout())

	client.builder.disableTimeouts = true

	assert.Equal(t, 0*time.Second, client.getResponseTimeout())

	client.builder.responseTimeout = 0 * time.Second
	client.builder.disableTimeouts = false

	assert.Equal(t, defaultResponseTimeout, client.getResponseTimeout())
}

func TestPrivateDo(t *testing.T) {
	// github.com/vpofe/go-http-client/gohttp/client_core.go:40: do 0.0%
}
