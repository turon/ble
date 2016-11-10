package dev

import (
	"github.com/currantlabs/ble"
	"github.com/currantlabs/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(id int) (d ble.Device, err error) {
	return darwin.NewDevice()
}
