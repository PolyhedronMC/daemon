package main

import (
	"github.com/polyhedronmc/daemon/daemon"
	"github.com/polyhedronmc/daemon/api"
	"github.com/polyhedronmc/daemon/config"
)

func main() {
	config := config.GetConfig()

	daemon := daemon.Create(config)
	server := api.Create(config)
	
	daemon.Start()
	server.Listen()
}
