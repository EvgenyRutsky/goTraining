package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient_ShouldNotReturnNil(t *testing.T) {
	c := NewClient()
	assert.NotNil(t, c)
}
