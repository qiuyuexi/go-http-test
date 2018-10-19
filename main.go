package main

import "go-test/pkg/server"

func main() {
	server.RegisterSignal()
	server.LogPid()
	server.Start()
}
