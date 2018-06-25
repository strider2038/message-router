package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/strider2038/message-router/config"
	"testing"
)

func TestNewDispatchingServer_Config_DispatchingServerCreated(t *testing.T) {
	server := NewDispatchingServer(config.Config{})

	assert.NotNil(t, server)
}
