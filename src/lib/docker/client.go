package docker

import (
	"github.com/docker/docker/client"
)

func GetClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}
