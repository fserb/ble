package dev

import (
	"github.com/melvinto/ble"
	"github.com/melvinto/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
