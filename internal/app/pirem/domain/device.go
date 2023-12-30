package domain

type Device struct {
	ID              DeviceID
	Name            string
	CanSend         bool
	CanReceive      bool
	DriverVersion   string
	FirmwareVersion string
}
