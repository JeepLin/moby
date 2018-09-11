package register // import "moby/daemon/graphdriver/register"

import (
	// register the windows graph drivers
	_ "moby/daemon/graphdriver/lcow"
	_ "moby/daemon/graphdriver/windows"
)
