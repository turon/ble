package dev

import "github.com/currantlabs/ble"

// NewDevice ...
func NewDevice(id int) (d ble.Device, err error) {
	return DefaultDevice(id)
}
