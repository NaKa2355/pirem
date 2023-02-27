package usecases

import (
	"context"

	irdatav1 "github.com/NaKa2355/irdeck-proto/gen/go/common/irdata/v1"
	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
)

type DeviceController interface {
	GetDeviceInfo(ctx context.Context) (*apiremv1.DeviceInfo, error)
	GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error)
	SendRawIr(ctx context.Context, ir_data *irdatav1.RawIrData) error
	ReceiveRawIr(ctx context.Context) (*irdatav1.RawIrData, error)
	IsBusy(ctx context.Context) (bool, error)
	Drop() error
}
