package repository

import "github.com/NaKa2355/pirem/internal/app/pirem/entity/device"

type DeviceCreator interface {
	CreateDevice(dev *device.Device) error
}

type DeviceReader interface {
	ReadDevice(id string) (*device.Device, error)
}

type DevicesReader interface {
	ReadDevices() ([]*device.Device, error)
}

type DeviceDeleter interface {
	DeleteDevice(id string) error
}
