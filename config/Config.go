package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Brokers []string
	Port    int
}

func LoadConfigFromEnvironment() Config {
	brokers := os.Getenv("MESSAGE_ROUTER_KAFKA_BROKERS")
	envPort := os.Getenv("MESSAGE_ROUTER_PORT")
	port, _ := strconv.ParseUint(envPort, 10, 32)

	if brokers == "" {
		brokers = "localhost:9092"
	}

	if port == 0 {
		port = 3000
	}

	return Config{
		strings.Split(brokers, ","),
		int(port),
	}
}
