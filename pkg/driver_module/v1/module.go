package driver_module

import "encoding/json"

type DriverModule interface {
	NewDriver(json.RawMessage) (Device, error)
}
