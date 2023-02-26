package usecases

import (
	"context"

	apiremv1 "github.com/NaKa2355/pirem/gen/apirem/v1"
)

type DeviceController interface {
	GetDeviceInfo(ctx context.Context) (*apiremv1.DeviceInfo, error)
	GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error)
	SendRawIr(ctx context.Context, ir_data *apiremv1.RawIrData) error
	ReceiveRawIr(ctx context.Context) (*apiremv1.RawIrData, error)
	IsBusy(ctx context.Context) (bool, error)
	Drop() error
}
