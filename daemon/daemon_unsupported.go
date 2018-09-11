// +build !linux,!freebsd,!windows

package daemon // import "moby/daemon"
import "moby/daemon/config"

const platformSupported = false

func setupResolvConf(config *config.Config) {
}
