// +build !exclude_graphdriver_zfs,linux !exclude_graphdriver_zfs,freebsd

package register // import "moby/daemon/graphdriver/register"

import (
	// register the zfs driver
	_ "moby/daemon/graphdriver/zfs"
)
