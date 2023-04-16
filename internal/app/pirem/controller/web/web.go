package web

import (
	"context"
	"errors"

	v1 "github.com/NaKa2355/irdeck-proto/gen/go/common/irdata/v1"
	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
)

type Interactor interface {
	bdy.IRSender
	bdy.IRReceiver
	bdy.DeviceInfoGetter
	bdy.DevicesInfoGetter
}

type Handler struct {
	apiremv1.UnimplementedPiRemServiceServer
	i Interactor
}

func New(interactor Interactor) *Handler {
	return &Handler{
		i: interactor,
	}
}

func (w *Handler) GetAllDeviceInfo(ctx context.Context, req *apiremv1.GetAllDeviceInfoRequest) (res *apiremv1.GetAllDeviceInfoResponse, err error) {
	res = &apiremv1.GetAllDeviceInfoResponse{}
	out, err := w.i.GetDevicesInfo(ctx)
	res.DeviceInfo = make([]*apiremv1.DeviceInfo, len(out.Devices))
	for i, d := range out.Devices {
		res.DeviceInfo[i] = &apiremv1.DeviceInfo{
			Id:              d.ID,
			Name:            d.Name,
			BufferSize:      int32(d.BufferSize),
			DriverVersion:   d.DriverVersion,
			FirmwareVersion: d.FirmwareVersion,
		}
	}
	return res, err
}

func (w *Handler) GetDeviceInfo(ctx context.Context, req *apiremv1.GetDeviceInfoRequest) (res *apiremv1.DeviceInfo, err error) {
	out, err := w.i.GetDeviceInfo(ctx, bdy.GetDeviceInput{ID: req.DeviceId})
	if err != nil {
		return res, err
	}
	dev := out.Device

	res = &apiremv1.DeviceInfo{
		Id:              dev.ID,
		Name:            dev.Name,
		BufferSize:      int32(dev.BufferSize),
		DriverVersion:   dev.DriverVersion,
		FirmwareVersion: dev.FirmwareVersion,
	}
	return res, err
}

func (w *Handler) ReceiveRawIr(ctx context.Context, req *apiremv1.ReceiveRawIrRequest) (res *apiremv1.ReceiveRawIrResponse, err error) {
	res = &apiremv1.ReceiveRawIrResponse{
		IrData: &v1.IrData{},
	}

	out, err := w.i.ReceiveIR(ctx, bdy.ReceiveIRInput{ID: req.DeviceId})
	if err != nil {
		return res, err
	}

	switch irdata := out.(type) {
	case bdy.RawIRData:
		res.IrData.Data = &v1.IrData_Raw{
			Raw: &v1.RawIrData{
				CarrierFreqKhz: irdata.CarrierFreqKiloHz,
				OnOffPluseNs:   irdata.PluseNanoSec,
			},
		}
	default:
		return res, errors.New("unsupported irdata")
	}

	return res, err
}
