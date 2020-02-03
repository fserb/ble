package dev

import (
	"github.com/melvinto/ble"
	"github.com/melvinto/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
