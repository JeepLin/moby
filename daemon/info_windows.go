package daemon // import "moby/daemon"

import (
	"moby/api/types"
	"moby/pkg/sysinfo"
)

// fillPlatformInfo fills the platform related info.
func (daemon *Daemon) fillPlatformInfo(v *types.Info, sysInfo *sysinfo.SysInfo) {
}

func fillDriverWarnings(v *types.Info) {
}
