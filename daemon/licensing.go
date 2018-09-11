package daemon // import "moby/daemon"

import (
	"moby/api/types"
	"moby/dockerversion"
)

func (daemon *Daemon) fillLicense(v *types.Info) {
	v.ProductLicense = dockerversion.DefaultProductLicense
}
