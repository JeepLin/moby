package container

import (
	"context"
	"testing"

	"moby/api/types"
	"moby/api/types/container"
	"moby/api/types/network"
	"moby/client"
	"gotest.tools/assert"
)

// TestContainerConfig holds container configuration struct that
// are used in api calls.
type TestContainerConfig struct {
	Name             string
	Config           *container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
}

// Create creates a container with the specified options
// nolint: golint
func Create(t *testing.T, ctx context.Context, client client.APIClient, ops ...func(*TestContainerConfig)) string { // nolint: golint
	t.Helper()
	config := &TestContainerConfig{
		Config: &container.Config{
			Image: "busybox",
			Cmd:   []string{"top"},
		},
		HostConfig:       &container.HostConfig{},
		NetworkingConfig: &network.NetworkingConfig{},
	}

	for _, op := range ops {
		op(config)
	}

	c, err := client.ContainerCreate(ctx, config.Config, config.HostConfig, config.NetworkingConfig, config.Name)
	assert.NilError(t, err)

	return c.ID
}

// Run creates and start a container with the specified options
// nolint: golint
func Run(t *testing.T, ctx context.Context, client client.APIClient, ops ...func(*TestContainerConfig)) string { // nolint: golint
	t.Helper()
	id := Create(t, ctx, client, ops...)

	err := client.ContainerStart(ctx, id, types.ContainerStartOptions{})
	assert.NilError(t, err)

	return id
}
