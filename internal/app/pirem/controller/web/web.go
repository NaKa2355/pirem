package web

import (
	"context"
	"errors"

	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
	v1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
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

var _ apiremv1.PiRemServiceServer = &Handler{}

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

func (w *Handler) SendIr(ctx context.Context, req *apiremv1.SendIrRequest) (res *apiremv1.SendIrResponse, err error) {
	res = &apiremv1.SendIrResponse{}
	switch irdata := req.IrData.Data.(type) {
	case *v1.IrData_Raw:
		err = w.i.SendIR(ctx, bdy.SendIRInput{
			ID: req.DeviceId,
			IRData: bdy.RawIRData{
				CarrierFreqKiloHz: irdata.Raw.GetCarrierFreqKhz(),
				PluseNanoSec:      irdata.Raw.GetOnOffPluseNs(),
			}})
	default:
		err = errors.New("unsupported error")
	}
	return res, err
}

func (w *Handler) ReceiveIr(ctx context.Context, req *apiremv1.ReceiveIrRequest) (res *apiremv1.ReceiveIrResponse, err error) {
	res = &apiremv1.ReceiveIrResponse{
		IrData: &v1.IrData{},
	}

	out, err := w.i.ReceiveIR(ctx, bdy.ReceiveIRInput{ID: req.DeviceId})
	if err != nil {
		return res, err
	}

	irdata := out.ConvertToRaw()
	res.IrData.Data = &v1.IrData_Raw{
		Raw: &v1.RawIrData{
			CarrierFreqKhz: irdata.CarrierFreqKiloHz,
			OnOffPluseNs:   irdata.PluseNanoSec,
		},
	}

	return res, err
}
