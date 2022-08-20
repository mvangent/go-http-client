package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	assert.NotNil(t, httpClient)
}
