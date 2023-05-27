package module

type DeviceInfo struct {
	CanSend         bool
	CanReceive      bool
	DriverVersion   string
	FirmwareVersion string
}
