// +build !linux

package daemon // import "moby/daemon"

func ensureDefaultAppArmorProfile() error {
	return nil
}
