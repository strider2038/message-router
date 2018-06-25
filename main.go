package main

import (
	config2 "bitbucket.org/strider2038/event-router/config"
	"bitbucket.org/strider2038/event-router/server"
)

func main() {
	config := config2.LoadConfigFromEnvironment()
	dispatchingServer := server.NewDispatchingServer(config)
	dispatchingServer.Start()
}
