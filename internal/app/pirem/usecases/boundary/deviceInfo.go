package boundary

type DeviceInfo struct {
	ID              string
	Name            string
	BufferSize      int
	CanSend         bool
	CanReceive      bool
	DriverVersion   string
	FirmwareVersion string
}
