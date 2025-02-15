package driver_module

import "encoding/json"

type DriverModule interface {
	LoadDevice(json.RawMessage) (Device, error)
}
