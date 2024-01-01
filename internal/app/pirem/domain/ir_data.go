package domain

type RawData struct {
	CarrierFreqKiloHz uint32
	PluseNanoSec      []uint32
}

func (r *RawData) ConvertToRaw() *RawData {
	return r
}

type IRData interface {
	ConvertToRaw() *RawData
}
