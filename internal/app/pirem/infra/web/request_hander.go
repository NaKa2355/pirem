package web

import (
	"context"

	api "github.com/NaKa2355/pirem/api/gen/go/api/v1"
	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/proto"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
)

type InputBoundary interface {
	bdy.ButtonGetter
	bdy.ButtonPusher

	bdy.RemoteCreator
	bdy.RemoteGetter
	bdy.RemoteLister
	bdy.RemoteUpdater
	bdy.RemoteDeleter
	bdy.IRLearner

	bdy.DeviceGetter
	bdy.DevicesLister
	bdy.IRSender
	bdy.IRReceiver
}

type RequestHander struct {
	port InputBoundary
	api.UnimplementedPiRemServiceServer
}

func (h *RequestHander) SendIr(ctx context.Context, req *api.SendIrRequest) (*api.SendIrResponse, error) {
	deviceID := req.GetDeviceId()
	err := h.port.SendIR(ctx, bdy.SendIRInput{
		ID:     deviceID,
		IRData: adapter.UnMarshalIRData(req.IrData),
	})
	return &api.SendIrResponse{}, err
}

func (h *RequestHander) ReceiveIr(ctx context.Context, req *api.ReceiveIrRequest) (*api.IrData, error) {
	deviceID := req.GetDeviceId()
	irData, err := h.port.ReceiveIR(ctx, bdy.ReceiveIRInput{
		ID: deviceID,
	})
	rawData := irData.ConvertToRaw()
	return &api.IrData{
		Data: &api.IrData_Raw{
			Raw: &api.RawIrData{
				CarrierFreqKhz: rawData.CarrierFreqKiloHz,
				OnOffPluseNs:   rawData.PluseNanoSec,
			},
		},
	}, err
}

func (h *RequestHander) ListDevices(ctx context.Context, req *api.ListDevicesRequest) (*api.ListDevicesResponse, error) {
	devices, err := h.port.ListDevices(ctx)
	devicesRes := []*api.Device{}
	for _, device := range devices {
		devicesRes = append(devicesRes, &api.Device{
			Id:              device.ID,
			Name:            device.Name,
			CanSend:         device.CanSend,
			CanReceive:      device.CanReceive,
			FirmwareVersion: device.FirmwareVersion,
			DriverVersion:   device.FirmwareVersion,
		})
	}
	return &api.ListDevicesResponse{
		Devices: devicesRes,
	}, err
}

func (h *RequestHander) GetDevice(ctx context.Context, req *api.GetDeviceRequest) (*api.Device, error) {
	device, err := h.port.GetDevice(ctx, bdy.GetDeivceInput{
		DeviceID: req.DeviceId,
	})
	return &api.Device{
		Id:              device.ID,
		Name:            device.Name,
		CanSend:         device.CanSend,
		CanReceive:      device.CanReceive,
		FirmwareVersion: device.FirmwareVersion,
		DriverVersion:   device.FirmwareVersion,
	}, err
}

// remotes
func (h *RequestHander) CreateRemote(ctx context.Context, req *api.Remote) (res *api.Remote, err error) {
	name, err := remote.NewName(req.Name)
	if err != nil {
		return
	}
	buttons := []struct {
		Name button.Name
		Tag  button.Tag
	}{}
	for _, createButtonsReq := range req.Buttons {
		name, err := button.NewName(createButtonsReq.Name)
		if err != nil {
			return res, err
		}
		buttons = append(buttons, struct {
			Name button.Name
			Tag  button.Tag
		}{
			Name: name,
			Tag:  button.Tag(createButtonsReq.Tag),
		})
	}

	remote, err := h.port.CreateRemote(ctx, bdy.CreateRemoteInput{
		Name:     name,
		Tag:      req.Tag,
		DeviveID: req.DeviceId,
		Buttons:  buttons,
	})
	req.Id = string(remote.ID)
	return req, err
}

func (h *RequestHander) ListRemotes(ctx context.Context, req *api.ListRemotesRequest) (res *api.ListRemotesResponse, err error) {
	remotes, err := h.port.ListRemotes(ctx)
	if err != nil {
		return
	}
	res = &api.ListRemotesResponse{}
	for _, remote := range remotes {
		buttonsRes := []*api.Remote_Button{}
		for _, button := range remote.Buttons {
			buttonsRes = append(buttonsRes, &api.Remote_Button{
				Id:        string(button.ID),
				Name:      string(button.Name),
				Tag:       string(button.Tag),
				HasIrData: button.IRData != nil,
			})
		}
		res.Remotes = append(res.Remotes, &api.Remote{
			Buttons: buttonsRes,
		})
	}
	return
}

func (h *RequestHander) GetRemote(ctx context.Context, req *api.GetRemoteRequest) (res *api.Remote, err error) {
	remote, err := h.port.GetRemote(ctx, bdy.GetRemoteInput{
		RemoteID: remote.ID(req.RemoteId),
	})
	if err != nil {
		return
	}

	buttonsRes := []*api.Remote_Button{}
	for _, button := range remote.Buttons {
		buttonsRes = append(buttonsRes, &api.Remote_Button{
			Id:        string(button.ID),
			Name:      string(button.Name),
			Tag:       string(button.Tag),
			HasIrData: button.IRData != nil,
		})
	}
	res = &api.Remote{
		Buttons: buttonsRes,
	}

	return
}

func (h *RequestHander) UpdateRemote(ctx context.Context, req *api.UpdateRemoteRequest) (res *api.Empty, err error) {
	name, err := remote.NewName(req.Name)
	if err != nil {
		return
	}
	err = h.port.UpdateRemote(ctx, bdy.UpdateRemoteInput{
		RemoteID:   remote.ID(req.RemoteId),
		RemoteName: name,
		DeviceID:   req.DeviceId,
	})
	return &api.Empty{}, err
}

func (h *RequestHander) DeleteRemote(ctx context.Context, req *api.DeleteRemoteRequest) (*api.Empty, error) {
	return &api.Empty{}, h.port.DeleteRemote(ctx, bdy.DeleteRemoteInput{
		RemoteID: remote.ID(req.RemoteId),
	})
}

// buttons
func (h *RequestHander) GetButton(ctx context.Context, req *api.GetButtonRequest) (res *api.Button, err error) {
	button, err := h.port.GetButton(ctx, bdy.GetButtonInput{
		ButtonID: button.ID(req.ButtonId),
	})
	if err != nil {
		return
	}
	return &api.Button{
		Id:     string(button.ID),
		Name:   string(button.Name),
		Tag:    string(button.Name),
		IrData: adapter.MarshalIRData(button.IRData),
	}, err
}

func (h *RequestHander) LearnIrData(ctx context.Context, req *api.LearnIrDataRequest) (*api.Empty, error) {
	return &api.Empty{}, h.port.LearnIR(ctx, bdy.LearnIRInput{
		ButtonID: button.ID(req.ButtonId),
		IRData:   adapter.UnMarshalIRData(req.IrData),
	})
}

func (h *RequestHander) PushButton(ctx context.Context, req *api.PushButtonRequest) (*api.Empty, error) {
	return &api.Empty{}, h.port.PushRemote(ctx, bdy.PushButtonInput{
		ButtonId: button.ID(req.ButtonId),
	})
}
