package device_controller

import (
	"context"
	apiremv1 "pirem/pkg/apirem.v1"
)

type DeviceController interface {
	GetDeviceInfo(ctx context.Context) (*apiremv1.DeviceInfo, error)
	GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error)
	SendIr(ctx context.Context, ir_data *apiremv1.RawIrData) error
	ReceiveIr(ctx context.Context) (*apiremv1.RawIrData, error)
	Drop() error
}
