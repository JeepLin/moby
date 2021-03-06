package container // import "moby/integration/container"

import (
	"context"
	"encoding/json"
	"io"
	"testing"
	"time"

	"moby/api/types"
	"moby/integration/internal/container"
	"moby/internal/test/request"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"gotest.tools/poll"
	"gotest.tools/skip"
)

func TestStats(t *testing.T) {
	skip.If(t, !testEnv.DaemonInfo.MemoryLimit)

	defer setupTest(t)()
	client := request.NewAPIClient(t)
	ctx := context.Background()

	info, err := client.Info(ctx)
	assert.NilError(t, err)

	cID := container.Run(t, ctx, client)

	poll.WaitOn(t, container.IsInState(ctx, client, cID, "running"), poll.WithDelay(100*time.Millisecond))

	resp, err := client.ContainerStats(ctx, cID, false)
	assert.NilError(t, err)
	defer resp.Body.Close()

	var v *types.Stats
	err = json.NewDecoder(resp.Body).Decode(&v)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(int64(v.MemoryStats.Limit), info.MemTotal))
	err = json.NewDecoder(resp.Body).Decode(&v)
	assert.Assert(t, is.ErrorContains(err, ""), io.EOF)
}
