package client // import "moby/client"

import (
	"context"
	"net/url"

	"moby/api/types"
)

// PluginRemove removes a plugin
func (cli *Client) PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error {
	query := url.Values{}
	if options.Force {
		query.Set("force", "1")
	}

	resp, err := cli.delete(ctx, "/plugins/"+name, query, nil)
	ensureReaderClosed(resp)
	return wrapResponseError(err, resp, "plugin", name)
}
