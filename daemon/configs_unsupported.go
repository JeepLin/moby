// +build !linux,!windows

package daemon // import "moby/daemon"

func configsSupported() bool {
	return false
}
