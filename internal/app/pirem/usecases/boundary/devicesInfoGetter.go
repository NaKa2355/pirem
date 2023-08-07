package boundary

import "context"

type GetDevicesInfoOutput struct {
	Devices []DeviceInfo
}

type DevicesInfoGetter interface {
	GetDevicesInfo(context.Context) (GetDevicesInfoOutput, error)
}
