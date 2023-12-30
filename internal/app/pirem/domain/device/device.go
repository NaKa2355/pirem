package device

type Device struct {
	ID              ID
	Name            string
	CanSend         bool
	CanReceive      bool
	DriverVersion   string
	FirmwareVersion string
}
