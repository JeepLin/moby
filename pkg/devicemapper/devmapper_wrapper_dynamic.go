// +build linux,cgo,!static_build

package devicemapper // import "moby/pkg/devicemapper"

// #cgo pkg-config: devmapper
import "C"
