package adapter

import (
	api "github.com/NaKa2355/pirem/api/gen/go/api/v1"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/irdata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type domainIRData struct {
	Data *api.IrData
}

func (data *domainIRData) ConvertToRaw() *irdata.RawData {
	return &irdata.RawData{}
}

func UnMarshalIRData(protoIRData *api.IrData) irdata.IRData {
	return &domainIRData{
		Data: protoIRData,
	}
}

func MarshalIRData(domainIRData irdata.IRData) *api.IrData {
	rawData := domainIRData.ConvertToRaw()
	return &api.IrData{
		Data: &api.IrData_Raw{
			Raw: &api.RawIrData{
				CarrierFreqKhz: rawData.CarrierFreqKiloHz,
				OnOffPluseNs:   rawData.PluseNanoSec,
			},
		},
	}
}

func UnmarshalBinaryIRData(binaryIRData []byte) (irdata.IRData, error) {
	irData := &anypb.Any{}
	err := proto.Unmarshal(binaryIRData, irData)
	if err != nil {
		return nil, err
	}
	protoIRData := &api.IrData{}
	err = anypb.UnmarshalTo(irData, protoIRData, proto.UnmarshalOptions{})
	return UnMarshalIRData(protoIRData), err
}

func MarshalIRDataToBinary(data irdata.IRData) ([]byte, error) {
	protoIRData := MarshalIRData(data)
	anyIRData := &anypb.Any{}
	err := anypb.MarshalFrom(anyIRData, protoIRData, proto.MarshalOptions{})
	if err != nil {
		return nil, err
	}
	return proto.Marshal(anyIRData)
}
