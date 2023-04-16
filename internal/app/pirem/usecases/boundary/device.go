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

type Status interface {
	status()
}

type StatuesWiredDevice struct {
	IsActive bool
}

func (w StatuesWiredDevice) status() {}

type IRData interface{}

type RawIRData struct {
	CarrierFreqKiloHz uint32
	PluseNanoSec      []uint32
}
