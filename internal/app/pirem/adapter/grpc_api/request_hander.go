package adapter

import (
	"context"

	api "github.com/NaKa2355/pirem-proto/gen/go/api/v1"
	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/proto"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/utils"
)

type RequestHander struct {
	port bdy.Boundary
	api.UnimplementedPiRemServiceServer
}

func NewRequestHandler(port bdy.Boundary) api.PiRemServiceServer {
	return &RequestHander{
		port: port,
	}
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
	return &api.ListDevicesResponse{
		Devices: utils.Map(devices, func(d *domain.Device) *api.Device {
			return &api.Device{
				Id:              d.ID,
				Name:            d.Name,
				CanSend:         d.CanSend,
				CanReceive:      d.CanReceive,
				FirmwareVersion: d.FirmwareVersion,
				DriverVersion:   d.FirmwareVersion,
			}
		}),
	}, err
}

func (h *RequestHander) GetDevice(ctx context.Context, req *api.GetDeviceRequest) (*api.Device, error) {
	d, err := h.port.GetDevice(ctx, bdy.GetDeivceInput{
		DeviceID: req.DeviceId,
	})
	return &api.Device{
		Id:              d.ID,
		Name:            d.Name,
		CanSend:         d.CanSend,
		CanReceive:      d.CanReceive,
		FirmwareVersion: d.FirmwareVersion,
		DriverVersion:   d.FirmwareVersion,
	}, err
}

// remotes
func (h *RequestHander) CreateRemote(ctx context.Context, req *api.CreateRemoteRequest) (res *api.Remote, err error) {
	name, err := domain.NewRemoteName(req.Name)
	if err != nil {
		return
	}
	buttons := []struct {
		Name domain.ButtonName
		Tag  domain.ButtonTag
	}{}
	for _, createButtonsReq := range req.Buttons {
		name, err := domain.NewButtonName(createButtonsReq.Name)
		if err != nil {
			return res, err
		}
		buttons = append(buttons, struct {
			Name domain.ButtonName
			Tag  domain.ButtonTag
		}{
			Name: name,
			Tag:  domain.ButtonTag(createButtonsReq.Tag),
		})
	}

	r, err := h.port.CreateRemote(ctx, bdy.CreateRemoteInput{
		Name:     name,
		Tag:      req.Tag,
		DeviveID: req.DeviceId,
		Buttons:  buttons,
	})

	return &api.Remote{
		Id:       string(r.ID),
		Name:     string(r.Name),
		DeviceId: r.DeviceID,
		Tag:      r.Tag,
		Buttons: utils.Map(r.Buttons,
			func(b *domain.Button) *api.Remote_Button {
				return &api.Remote_Button{
					Id:        string(b.ID),
					Name:      string(b.Name),
					Tag:       string(b.Tag),
					HasIrData: b.HasIRData,
				}
			},
		),
	}, err
}

func (h *RequestHander) ListRemotes(ctx context.Context, req *api.ListRemotesRequest) (res *api.ListRemotesResponse, err error) {
	remotes, err := h.port.ListRemotes(ctx)
	if err != nil {
		return
	}
	return &api.ListRemotesResponse{
		Remotes: utils.Map(remotes,
			func(r *domain.Remote) *api.Remote {
				return &api.Remote{
					Id:       string(r.ID),
					Name:     string(r.Name),
					Tag:      r.Tag,
					DeviceId: r.DeviceID,
					Buttons: utils.Map(r.Buttons,
						func(b *domain.Button) *api.Remote_Button {
							return &api.Remote_Button{
								Id:        string(b.ID),
								Name:      string(b.Name),
								Tag:       string(b.Tag),
								HasIrData: b.HasIRData,
							}
						},
					),
				}
			},
		),
	}, nil
}

func (h *RequestHander) GetRemote(ctx context.Context, req *api.GetRemoteRequest) (res *api.Remote, err error) {
	r, err := h.port.GetRemote(ctx, bdy.GetRemoteInput{
		RemoteID: domain.RemoteID(req.RemoteId),
	})
	if err != nil {
		return
	}

	return &api.Remote{
		Id:       string(r.ID),
		Tag:      r.Tag,
		Name:     string(r.Name),
		DeviceId: r.DeviceID,
		Buttons: utils.Map(r.Buttons,
			func(b *domain.Button) *api.Remote_Button {
				return &api.Remote_Button{
					Id:   string(b.ID),
					Name: string(b.Name),
					Tag:  string(b.Tag),
				}
			},
		),
	}, nil
}

func (h *RequestHander) UpdateRemote(ctx context.Context, req *api.UpdateRemoteRequest) (res *api.Empty, err error) {
	name, err := domain.NewRemoteName(req.Name)
	if err != nil {
		return
	}
	err = h.port.UpdateRemote(ctx, bdy.UpdateRemoteInput{
		RemoteID:   domain.RemoteID(req.RemoteId),
		RemoteName: name,
		DeviceID:   req.DeviceId,
	})
	return &api.Empty{}, err
}

func (h *RequestHander) DeleteRemote(ctx context.Context, req *api.DeleteRemoteRequest) (*api.Empty, error) {
	return &api.Empty{}, h.port.DeleteRemote(ctx, bdy.DeleteRemoteInput{
		RemoteID: domain.RemoteID(req.RemoteId),
	})
}

// buttons
func (h *RequestHander) GetButton(ctx context.Context, req *api.GetButtonRequest) (res *api.Button, err error) {
	b, err := h.port.GetButton(ctx, bdy.GetButtonInput{
		ButtonID: domain.ButtonID(req.ButtonId),
	})
	if err != nil {
		return
	}
	return &api.Button{
		Id:        string(b.ID),
		Name:      string(b.Name),
		Tag:       string(b.Name),
		HasIrData: b.HasIRData,
	}, err
}

func (h *RequestHander) LearnIrData(ctx context.Context, req *api.LearnIrDataRequest) (*api.Empty, error) {
	return &api.Empty{}, h.port.LearnIR(ctx, bdy.LearnIRInput{
		ButtonID: domain.ButtonID(req.ButtonId),
		IRData:   adapter.UnMarshalIRData(req.IrData),
	})
}

func (h *RequestHander) PushButton(ctx context.Context, req *api.PushButtonRequest) (*api.Empty, error) {
	return &api.Empty{}, h.port.PushRemote(ctx, bdy.PushButtonInput{
		ButtonId: domain.ButtonID(req.ButtonId),
	})
}

func (h *RequestHander) GetIrData(ctx context.Context, req *api.GetIrDataRequest) (*api.GetIrDataResponse, error) {
	out, err := h.port.GetIR(ctx, bdy.GetIRInput{ButtonID: domain.ButtonID(req.ButtonId)})
	return &api.GetIrDataResponse{
		IrData:   adapter.MarshalIRData(out.IRData),
		DeviceId: out.DeviceID,
	}, err
}
