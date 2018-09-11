// +build !exclude_graphdriver_devicemapper,!static_build,linux

package register // import "moby/daemon/graphdriver/register"

import (
	// register the devmapper graphdriver
	_ "moby/daemon/graphdriver/devmapper"
)
