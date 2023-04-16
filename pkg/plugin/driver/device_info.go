package driver

type ServiceType int

type DeviceInfo struct {
	CanSend         bool
	CanReceive      bool
	DriverVersion   string
	FirmwareVersion string
}
