package daemon // import "moby/daemon"

import (
	"testing"

	"moby/api/types"
	"moby/dockerversion"
	"gotest.tools/assert"
)

func TestfillLicense(t *testing.T) {
	v := &types.Info{}
	d := &Daemon{
		root: "/var/lib/docker/",
	}
	d.fillLicense(v)
	assert.Assert(t, v.ProductLicense == dockerversion.DefaultProductLicense)
}
