package docker

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/polyhedronmc/daemon/utils"
)

var docker *client.Client
var ctx = context.Background()
var connected bool = false


// check if docker is alive
func isDockerAlive() bool {
	if (docker == nil || !connected) {
		return false
	}
	return true
}

// SetupClient Set up the Docker client.
func SetupClient() {
	utils.PanicIf(isDockerAlive(), "docker already alive")

	var err error
	docker, err = client.NewEnvClient()
	if (err != nil) {
		panic(err)
	}
	connected = true
}

// GetVersion Return the current version of the host Docker engine.
func GetVersion() string {
	utils.PanicIf(!isDockerAlive(), "docker not alive")

	version, err := docker.ServerVersion(ctx)
	if (err != nil) {
		panic(err)
	}

	return version.Version
}

// GetContainer Get the container with the specified name.
func GetContainer(container string) {
	utils.PanicIf(!isDockerAlive(), "docker not alive")
}	
