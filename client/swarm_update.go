package client // import "moby/client"

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"moby/api/types/swarm"
)

// SwarmUpdate updates the swarm.
func (cli *Client) SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error {
	query := url.Values{}
	query.Set("version", strconv.FormatUint(version.Index, 10))
	query.Set("rotateWorkerToken", fmt.Sprintf("%v", flags.RotateWorkerToken))
	query.Set("rotateManagerToken", fmt.Sprintf("%v", flags.RotateManagerToken))
	query.Set("rotateManagerUnlockKey", fmt.Sprintf("%v", flags.RotateManagerUnlockKey))
	resp, err := cli.post(ctx, "/swarm/update", query, swarm, nil)
	ensureReaderClosed(resp)
	return err
}
