package gohttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientBuilder(t *testing.T) {
	builder := NewBuilder()

	assert.NotNil(t, builder)
}
