package driver_module

import "context"

// driver must impliment "GetDriver(json.RawMessage) (Device, Error)"
type Device interface {
	// SendIR(context.Context, *IRData) error
	// ReceiveIR(context.Context) (*IRData, error)
	GetInfo(context.Context) (*DeviceInfo, error)
	Drop() error
}

type Sender interface {
	SendIR(context.Context, *IRData) error
}

type Receiver interface {
	ReceiveIR(context.Context) (*IRData, error)
}
