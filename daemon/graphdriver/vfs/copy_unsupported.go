// +build !linux

package vfs // import "moby/daemon/graphdriver/vfs"

import "moby/pkg/chrootarchive"

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(nil).CopyWithTar(srcDir, dstDir)
}
