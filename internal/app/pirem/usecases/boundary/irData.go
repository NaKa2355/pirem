package boundary

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
