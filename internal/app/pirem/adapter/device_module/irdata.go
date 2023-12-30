package adapter

import (
	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	"github.com/NaKa2355/pirem/pkg/module/v1"
)

type domainIRData struct {
	irData *module.IRData
}

func (i *domainIRData) ConvertToRaw() *domain.RawData {
	return &domain.RawData{
		CarrierFreqKiloHz: i.irData.CarrierFreqKiloHz,
		PluseNanoSec:      i.irData.PluseNanoSec,
	}
}

func MarshalIRData(from domain.IRData) *module.IRData {
	rawData := from.ConvertToRaw()
	return &module.IRData{
		CarrierFreqKiloHz: rawData.CarrierFreqKiloHz,
		PluseNanoSec:      rawData.PluseNanoSec,
	}
}

func UnMarshalIRData(from *module.IRData) domain.IRData {
	return &domainIRData{irData: from}
}
