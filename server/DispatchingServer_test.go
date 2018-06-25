package server

import (
	"bitbucket.org/strider2038/event-router/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDispatchingServer_Config_DispatchingServerCreated(t *testing.T) {
	server := NewDispatchingServer(config.Config{})

	assert.NotNil(t, server)
}
