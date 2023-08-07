package boundary

import (
	"context"
	"encoding/json"

	"github.com/NaKa2355/pirem/pkg/module/v1"
)

type AddDeviceInput struct {
	ID         string
	Module     module.Module
	DeviceName string
	Config     json.RawMessage
}

type DeviceAdder interface {
	AddDevice(context.Context, AddDeviceInput) error
}
