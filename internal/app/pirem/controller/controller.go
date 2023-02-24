package controller

import (
	"context"
	"pirem/internal/app/pirem/usecases"

	apiremv1 "github.com/NaKa2355/pirem_pkg/apirem.v1"
)

type Controller struct {
	modelUsecase usecases.EntityController
	apiremv1.UnimplementedPiRemServiceServer
}

func New(entityUsecase usecases.EntityController) *Controller {
	return &Controller{modelUsecase: entityUsecase}
}

func (c *Controller) SendRawIr(ctx context.Context, in *apiremv1.SendRawIrRequest,
) (*apiremv1.SendRawIrResponse, error) {
	res := &apiremv1.SendRawIrResponse{}
	id := in.DeviceId
	err := c.modelUsecase.SendRawIr(ctx, id, in.IrData)
	return res, err
}

func (c *Controller) ReceiveRawIr(ctx context.Context, in *apiremv1.ReceiveRawIrRequest,
) (*apiremv1.ReceiveRawIrResponse, error) {
	res := &apiremv1.ReceiveRawIrResponse{}
	id := in.DeviceId

	irData, err := c.modelUsecase.ReceiveRawIr(ctx, id)
	if err != nil {
		return res, err
	}

	res.IrData = irData
	return res, err
}

func (c *Controller) GetAllDeviceInfo(ctx context.Context, in *apiremv1.GetAllDeviceInfoRequest,
) (*apiremv1.GetAllDeviceInfoResponse, error) {
	res := &apiremv1.GetAllDeviceInfoResponse{}

	info, err := c.modelUsecase.GetAllDeviceInfo(ctx)
	if err != nil {
		return res, err
	}

	res.DeviceInfo = info
	return res, err
}

func (c *Controller) GetDeviceInfo(ctx context.Context, in *apiremv1.GetDeviceInfoRequest,
) (*apiremv1.GetDeviceInfoResponse, error) {
	id := in.DeviceId
	res := &apiremv1.GetDeviceInfoResponse{}

	info, err := c.modelUsecase.GetDeviceInfo(ctx, id)
	if err != nil {
		return res, err
	}

	res.DeviceInfo = info
	return res, err
}

func (c *Controller) GetDeviceStatus(ctx context.Context, in *apiremv1.GetDeviceStatusRequest,
) (*apiremv1.GetDeviceStatusResponse, error) {
	id := in.DeviceId
	res := &apiremv1.GetDeviceStatusResponse{}

	status, err := c.modelUsecase.GetDeviceStatus(ctx, id)
	if err != nil {
		return res, err
	}

	res.DeviceStatus = status
	return res, err
}

func (c *Controller) IsBusy(ctx context.Context, in *apiremv1.IsBusyRequest,
) (*apiremv1.IsBusyResponse, error) {
	res := &apiremv1.IsBusyResponse{}
	id := in.DeviceId

	isBusy, err := c.modelUsecase.IsBusy(ctx, id)
	if err != nil {
		return res, err
	}

	res.IsBusy = isBusy
	return res, err
}
