package adapter

import (
	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	"github.com/NaKa2355/pirem/pkg/driver_module/v1"
)

type domainIRData struct {
	irData *driver_module.IRData
}

func (i *domainIRData) ConvertToRaw() *domain.RawData {
	return &domain.RawData{
		CarrierFreqKiloHz: i.irData.CarrierFreqKiloHz,
		PluseNanoSec:      i.irData.PluseNanoSec,
	}
}

func MarshalIRData(from domain.IRData) *driver_module.IRData {
	rawData := from.ConvertToRaw()
	return &driver_module.IRData{
		CarrierFreqKiloHz: rawData.CarrierFreqKiloHz,
		PluseNanoSec:      rawData.PluseNanoSec,
	}
}

func UnMarshalIRData(from *driver_module.IRData) domain.IRData {
	return &domainIRData{irData: from}
}
