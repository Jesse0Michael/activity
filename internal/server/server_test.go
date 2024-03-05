package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	server := New(Config{}, nil)

	assert.NotNil(t, server.Server, "router should not be nil")
}
