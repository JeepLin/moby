// +build !linux !cgo static_build !journald

package journald // import "moby/daemon/logger/journald"

func (s *journald) Close() error {
	return nil
}
