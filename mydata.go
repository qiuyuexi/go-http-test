package main

import "mydata/pkg/server"

func main() {
	server.RegisterSignal()
	server.LogPid()
	server.Start()
}
