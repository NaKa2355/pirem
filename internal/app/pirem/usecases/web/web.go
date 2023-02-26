package usecases

import (
	"context"
	"encoding/json"

	apiremv1 "github.com/NaKa2355/pirem/gen/apirem/v1"
)

type WebBoundary interface {
	GetAllDeviceInfo(ctx context.Context) ([]*apiremv1.DeviceInfo, error)
	GetDeviceInfo(ctx context.Context, id string) (*apiremv1.DeviceInfo, error)
	GetDeviceStatus(ctx context.Context, id string) (*apiremv1.DeviceStatus, error)
	IsBusy(ctx context.Context, id string) (bool, error)
	SendRawIr(ctx context.Context, id string, ir_data *apiremv1.RawIrData) error
	ReceiveRawIr(ctx context.Context, id string) (*apiremv1.RawIrData, error)
	AddDevice(id string, name string, pluginFilePath string, conf json.RawMessage) error
}
