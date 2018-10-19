package main

import "go-test/pkg/server"

func main() {
	server.LogInit()
	server.RegisterSignal()
	server.LogPid()
	server.Start()
}
