package driver

// driver must impliment "GetDriver(json.RawMessage) (Driver, Error)"
type Driver interface {
	SendIR(IRData) error
	ReceiveIR() (IRData, error)
	GetInfo() (DeviceInfo, error)
}
