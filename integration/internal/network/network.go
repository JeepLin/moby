package network

import (
	"context"
	"testing"

	"moby/api/types"
	"moby/client"
	"gotest.tools/assert"
)

func createNetwork(ctx context.Context, client client.APIClient, name string, ops ...func(*types.NetworkCreate)) (string, error) {
	config := types.NetworkCreate{}

	for _, op := range ops {
		op(&config)
	}

	n, err := client.NetworkCreate(ctx, name, config)
	return n.ID, err
}

// Create creates a network with the specified options
func Create(ctx context.Context, client client.APIClient, name string, ops ...func(*types.NetworkCreate)) (string, error) {
	return createNetwork(ctx, client, name, ops...)
}

// CreateNoError creates a network with the specified options and verifies there were no errors
// nolint: golint
func CreateNoError(t *testing.T, ctx context.Context, client client.APIClient, name string, ops ...func(*types.NetworkCreate)) string { // nolint: golint
	t.Helper()

	name, err := createNetwork(ctx, client, name, ops...)
	assert.NilError(t, err)
	return name
}
