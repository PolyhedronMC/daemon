package docker

import (
	"context"
	"github.com/polyhedronmc/daemon/config"
	"github.com/docker/docker/client"
)

// DetectDocker Check if Docker is running on the configuration declared.
func DetectDocker(config config.DaemonConfig) string {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	
	version, err := cli.ServerVersion(ctx)
	if (err != nil) {
		panic(err)
	}

	return version.Version
}
