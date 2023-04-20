package plugin

import "context"

// driver must impliment "GetDriver(json.RawMessage) (Driver, Error)"
type Driver interface {
	SendIR(context.Context, *IRData) error
	ReceiveIR(context.Context) (*IRData, error)
	GetInfo(context.Context) (*DeviceInfo, error)
	Drop() error
}
