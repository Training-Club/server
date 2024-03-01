package main

import (
	"tc-server/config"
	"tc-server/server"
)

func main() {
	conf := config.GetConfig()
	server.Init(conf)
}
