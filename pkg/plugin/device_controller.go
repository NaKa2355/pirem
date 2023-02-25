package plugin

import (
	"context"
	"encoding/json"

	apiremv1 "github.com/NaKa2355/pirem/pkg/apirem.v1"
)

type DeviceController interface {
	SendRawIr(context.Context, *apiremv1.RawIrData) error
	ReceiveRawIr(context.Context) (*apiremv1.RawIrData, error)
	GetDeviceInfo(context.Context) (*apiremv1.DeviceInfo, error)
	GetDeviceStatus(context.Context) (*apiremv1.DeviceStatus, error)
	Init(context.Context, json.RawMessage) error
}
