package main

import "bitbucket.org/strider2038/event-router/server"

func main() {
	dispatchingServer := server.NewDispatchingServer()
	dispatchingServer.Start()
}
