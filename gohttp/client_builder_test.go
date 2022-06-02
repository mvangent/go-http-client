package gohttp

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientBuilder(t *testing.T) {
	builder := NewBuilder()

	assert.NotNil(t, builder)

	var (
		disableTimeouts    = true
		connectionTimeout  = 2 * time.Second
		responseTimeout    = 12 * time.Second
		headers            = make(http.Header)
		userAgent          = "go-builder-test"
		maxIdleConnections = 12
		customClient       = http.Client{Timeout: 25 * time.Second}
	)

	headers.Set("X-Test_Builder", "True")

	builderConfig := builder.
		DisableTimeouts(disableTimeouts).
		SetConnectionTimeout(connectionTimeout).
		SetResponseTimeout(responseTimeout).
		SetHeaders(headers).
		SetUserAgent(userAgent).
		SetMaxIdleConnections(maxIdleConnections).
		SetHttpClient(&customClient).
		ReadBuilderConfig()

	assert.Equal(t, builderConfig.DisableTimeouts, disableTimeouts)
	assert.Equal(t, builderConfig.ConnectionTimeout, connectionTimeout)
	assert.Equal(t, builderConfig.ResponseTimeout, responseTimeout)
	assert.Equal(t, builderConfig.Headers, headers)
	assert.Equal(t, builderConfig.UserAgent, userAgent)
	assert.Equal(t, builderConfig.MaxIdleConnections, maxIdleConnections)
	assert.Equal(t, builderConfig.CustomClient, customClient)
}
