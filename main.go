package main

import (
	config2 "github.com/strider2038/message-router/config"
	"github.com/strider2038/message-router/server"
)

func main() {
	config := config2.LoadConfigFromEnvironment()
	dispatchingServer := server.NewDispatchingServer(config)
	dispatchingServer.Start()
}
