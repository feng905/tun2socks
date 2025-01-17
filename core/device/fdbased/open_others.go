//go:build !(linux && amd64) && !(linux && arm64) && !windows

package fdbased

import (
	"fmt"
	"os"

	"github.com/feng905/tun2socks/v2/core/device"
	"github.com/feng905/tun2socks/v2/core/device/iobased"
)

func open(fd int, mtu uint32, offset int) (device.Device, error) {
	f := &FD{fd: fd, mtu: mtu}

	ep, err := iobased.New(os.NewFile(uintptr(fd), f.Name()), mtu, offset)
	if err != nil {
		return nil, fmt.Errorf("create endpoint: %w", err)
	}
	f.LinkEndpoint = ep

	return f, nil
}
