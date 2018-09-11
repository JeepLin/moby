// +build !linux,!darwin,!freebsd,!windows

package daemon // import "moby/daemon"

func (d *Daemon) setupDumpStackTrap(_ string) {
	return
}
