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

type IRData interface {
	ConvertToRaw() RawIRData
}

type RawIRData struct {
	CarrierFreqKiloHz uint32
	PluseNanoSec      []uint32
}

func (data RawIRData) ConvertToRaw() RawIRData {
	return data
}
