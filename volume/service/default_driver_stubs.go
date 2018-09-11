// +build !linux,!windows

package service // import "moby/volume/service"

import (
	"moby/pkg/idtools"
	"moby/volume/drivers"
)

func setupDefaultDriver(_ *drivers.Store, _ string, _ idtools.Identity) error { return nil }
