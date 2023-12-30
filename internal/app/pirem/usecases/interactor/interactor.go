package interactor

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/irdata"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	gateway "github.com/NaKa2355/pirem/internal/app/pirem/usecases/gateways"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/irdevice"
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

type Interactor struct {
	repo   gateway.Repository
	device irdevice.IRDevice
}

var _ InputBoundary = &Interactor{}

func NewInteractor(repo gateway.Repository, device irdevice.IRDevice) *Interactor {
	return &Interactor{
		repo:   repo,
		device: device,
	}
}

func (i *Interactor) GetButton(ctx context.Context, input bdy.GetButtonInput) (*button.Button, error) {
	return i.repo.ReadButton(ctx, input.ButtonID)
}

func (i *Interactor) PushRemote(ctx context.Context, input bdy.PushButtonInput) error {
	button, err := i.repo.ReadButton(ctx, input.ButtonId)
	if err != nil {
		return err
	}
	return i.device.SendIR(ctx, button.DeviceID, button.IRData)
}

func (i *Interactor) CreateRemote(ctx context.Context, input bdy.CreateRemoteInput) (*remote.Remote, error) {
	buttons := []*button.Button{}
	for _, b := range input.Buttons {
		buttons = append(buttons, button.Factory(b.Name, b.Tag))
	}
	return i.repo.CreateRemote(
		ctx,
		remote.RemoteFactory(input.Name, input.DeviveID, input.Tag, buttons),
	)
}

func (i *Interactor) GetRemote(ctx context.Context, input bdy.GetRemoteInput) (*remote.Remote, error) {
	return i.repo.ReadRemote(ctx, input.RemoteID)
}

func (i *Interactor) ListRemotes(ctx context.Context) ([]*remote.Remote, error) {
	return i.repo.ReadRemotes(ctx)
}

func (i *Interactor) UpdateRemote(ctx context.Context, input bdy.UpdateRemoteInput) error {
	remote, err := i.repo.ReadRemote(ctx, input.RemoteID)
	if err != nil {
		return err
	}
	remote.UpdateRemote(input.RemoteName, input.DeviceID)
	return i.repo.UpdateRemote(ctx, remote)
}

func (i *Interactor) DeleteRemote(ctx context.Context, input bdy.DeleteRemoteInput) error {
	return i.repo.DeleteRemote(ctx, input.RemoteID)
}

func (i *Interactor) LearnIR(ctx context.Context, input bdy.LearnIRInput) error {
	button, err := i.repo.ReadButton(ctx, input.ButtonID)
	if err != nil {
		return err
	}
	button.LearnIR(input.IRData)
	return i.repo.UpdateButton(ctx, button)
}

func (i *Interactor) GetDevice(ctx context.Context, input bdy.GetDeivceInput) (*device.Device, error) {
	return i.device.ReadDevice(ctx, input.DeviceID)
}

func (i *Interactor) ListDevices(ctx context.Context) ([]*device.Device, error) {
	return i.device.ReadDevices(ctx)
}

func (i *Interactor) SendIR(ctx context.Context, input bdy.SendIRInput) error {
	return i.device.SendIR(ctx, input.ID, input.IRData)
}

func (i *Interactor) ReceiveIR(ctx context.Context, input bdy.ReceiveIRInput) (irdata.IRData, error) {
	return i.device.ReceiveIR(ctx, input.ID)
}
