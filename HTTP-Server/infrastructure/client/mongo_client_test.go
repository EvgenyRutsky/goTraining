package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient_ShouldNotReturnNil(t *testing.T) {
	config := NewConfig()
	c := NewClient(config)
	assert.NotNil(t, c)
}
