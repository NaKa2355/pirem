package web

import (
	"context"

	web_usecases "github.com/NaKa2355/pirem/internal/app/pirem/usecases/web"

	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
)

type Web struct {
	webBoundary web_usecases.WebBoundary
	apiremv1.UnimplementedPiRemServiceServer
}

func New(webBoundary web_usecases.WebBoundary) *Web {
	return &Web{webBoundary: webBoundary}
}

func (w *Web) SendRawIr(ctx context.Context, in *apiremv1.SendRawIrRequest,
) (*apiremv1.SendRawIrResponse, error) {
	res := &apiremv1.SendRawIrResponse{}
	id := in.DeviceId
	err := w.webBoundary.SendRawIr(ctx, id, in.IrData)
	return res, err
}

func (w *Web) ReceiveRawIr(ctx context.Context, in *apiremv1.ReceiveRawIrRequest,
) (*apiremv1.ReceiveRawIrResponse, error) {
	res := &apiremv1.ReceiveRawIrResponse{}
	id := in.DeviceId

	irData, err := w.webBoundary.ReceiveRawIr(ctx, id)
	if err != nil {
		return res, err
	}

	res.IrData = irData
	return res, err
}

func (w *Web) GetAllDeviceInfo(ctx context.Context, in *apiremv1.GetAllDeviceInfoRequest,
) (*apiremv1.GetAllDeviceInfoResponse, error) {
	res := &apiremv1.GetAllDeviceInfoResponse{}

	info, err := w.webBoundary.GetAllDeviceInfo(ctx)
	if err != nil {
		return res, err
	}

	res.DeviceInfo = info
	return res, err
}

func (w *Web) GetDeviceInfo(ctx context.Context, in *apiremv1.GetDeviceInfoRequest,
) (*apiremv1.GetDeviceInfoResponse, error) {
	id := in.DeviceId
	res := &apiremv1.GetDeviceInfoResponse{}

	info, err := w.webBoundary.GetDeviceInfo(ctx, id)
	if err != nil {
		return res, err
	}

	res.DeviceInfo = info
	return res, err
}

func (w *Web) GetDeviceStatus(ctx context.Context, in *apiremv1.GetDeviceStatusRequest,
) (*apiremv1.GetDeviceStatusResponse, error) {
	id := in.DeviceId
	res := &apiremv1.GetDeviceStatusResponse{}

	status, err := w.webBoundary.GetDeviceStatus(ctx, id)
	if err != nil {
		return res, err
	}

	res.DeviceStatus = status
	return res, err
}

func (w *Web) IsBusy(ctx context.Context, in *apiremv1.IsBusyRequest,
) (*apiremv1.IsBusyResponse, error) {
	res := &apiremv1.IsBusyResponse{}
	id := in.DeviceId

	isBusy, err := w.webBoundary.IsBusy(ctx, id)
	if err != nil {
		return res, err
	}

	res.IsBusy = isBusy
	return res, err
}
