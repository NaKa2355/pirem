package boundary

import "context"

type GetDeviceInput struct {
	ID string
}

type GetDeviceInfoOutput struct {
	Device DeviceInfo
}

type DeviceInfoGetter interface {
	GetDeviceInfo(context.Context, GetDeviceInput) (GetDeviceInfoOutput, error)
}
