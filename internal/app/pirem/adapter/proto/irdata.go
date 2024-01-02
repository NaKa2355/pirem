package adapter

import (
	api "github.com/NaKa2355/pirem/api/gen/go/api/v1"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type domainIRData struct {
	Data *api.IrData
}

func (data *domainIRData) ConvertToRaw() *domain.RawData {
	switch irData := data.Data.Data.(type) {
	case *api.IrData_Raw:
		return &domain.RawData{
			CarrierFreqKiloHz: irData.Raw.CarrierFreqKhz,
			PluseNanoSec:      irData.Raw.OnOffPluseNs,
		}
	default:
		return &domain.RawData{}
	}
}

func UnMarshalIRData(from *api.IrData) domain.IRData {
	return &domainIRData{
		Data: from,
	}
}

func MarshalIRData(from domain.IRData) *api.IrData {
	rawData := from.ConvertToRaw()
	return &api.IrData{
		Data: &api.IrData_Raw{
			Raw: &api.RawIrData{
				CarrierFreqKhz: rawData.CarrierFreqKiloHz,
				OnOffPluseNs:   rawData.PluseNanoSec,
			},
		},
	}
}

func UnmarshalBinaryIRData(from []byte) (domain.IRData, error) {
	data := &anypb.Any{}
	err := proto.Unmarshal(from, data)
	if err != nil {
		return &domain.RawData{}, err
	}
	protoIRData := &api.IrData{}
	err = anypb.UnmarshalTo(data, protoIRData, proto.UnmarshalOptions{})
	return UnMarshalIRData(protoIRData), err
}

func MarshalIRDataToBinary(from domain.IRData) ([]byte, error) {
	data := &anypb.Any{}
	err := anypb.MarshalFrom(data, MarshalIRData(from), proto.MarshalOptions{})
	if err != nil {
		return []byte{}, err
	}
	return proto.Marshal(data)
}
