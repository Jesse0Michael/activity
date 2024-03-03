package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	server := New(Config{})

	assert.NotNil(t, server.router, "router should not be nil")
}
