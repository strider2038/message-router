package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromEnvironment_AllParametersInEnv_ParametersLoaded(t *testing.T) {
	os.Setenv("MESSAGE_ROUTER_KAFKA_BROKERS", "localhost:9092,localhost:9093")
	os.Setenv("MESSAGE_ROUTER_PORT", "5000")

	config := LoadConfigFromEnvironment()

	assert.Equal(t, "localhost:9092", config.Brokers[0])
	assert.Equal(t, "localhost:9093", config.Brokers[1])
	assert.Equal(t, 5000, config.Port)
}

func TestLoadConfigFromEnvironment_NoParametersInEnv_DefaultParametersLoaded(t *testing.T) {
	os.Unsetenv("MESSAGE_ROUTER_KAFKA_BROKERS")
	os.Unsetenv("MESSAGE_ROUTER_PORT")

	config := LoadConfigFromEnvironment()

	assert.Equal(t, "localhost:9092", config.Brokers[0])
	assert.Equal(t, 3000, config.Port)
}
