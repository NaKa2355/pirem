package boundary

import "context"

type IRSender interface {
	SendIR(context.Context, SendIRInput) error
}

type IRReceiver interface {
	ReceiveIR(context.Context, ReceiveIRInput) (IRData, error)
}

type DevicesInfoGetter interface {
	GetDevicesInfo(context.Context) (GetDevicesInfoOutput, error)
}

type DeviceInfoGetter interface {
	GetDeviceInfo(context.Context, GetDeviceInput) (GetDeviceInfoOutput, error)
}
