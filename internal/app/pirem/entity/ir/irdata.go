package ir

type RawData struct {
	CarrierFreqKiloHz uint32
	PluseNanoSec      []uint32
}

func (r RawData) ConvertToRaw() RawData {
	return r
}

type Data interface {
	ConvertToRaw() RawData
}
