package device_plugin

import (
	"context"
	"encoding/json"

	irdatav1 "github.com/NaKa2355/irdeck-proto/gen/go/common/irdata/v1"
	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
)

type DeviceController interface {
	SendRawIr(context.Context, *irdatav1.RawIrData) error
	ReceiveRawIr(context.Context) (*irdatav1.RawIrData, error)
	GetDeviceInfo(context.Context) (*apiremv1.DeviceInfo, error)
	GetDeviceStatus(context.Context) (*apiremv1.DeviceStatus, error)
	IsBusy(context.Context) (bool, error)
	Init(context.Context, json.RawMessage) error
}
