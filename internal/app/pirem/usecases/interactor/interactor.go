package interactor

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	gateway "github.com/NaKa2355/pirem/internal/app/pirem/usecases/gateways"
)

type Interactor struct {
	repo   gateway.Repository
	device controllers.IRDevice
}

var _ bdy.Boundary = &Interactor{}

func NewInteractor(repo gateway.Repository, device controllers.IRDevice) *Interactor {
	return &Interactor{
		repo:   repo,
		device: device,
	}
}

func (i *Interactor) GetButton(ctx context.Context, input bdy.GetButtonInput) (*domain.Button, error) {
	return i.repo.ReadButton(ctx, input.ButtonID)
}

func (i *Interactor) PushRemote(ctx context.Context, input bdy.PushButtonInput) error {
	b, err := i.repo.ReadButton(ctx, input.ButtonId)
	if err != nil {
		return err
	}
	return i.device.SendIR(ctx, b.DeviceID, b.IRData)
}

func (i *Interactor) CreateRemote(ctx context.Context, input bdy.CreateRemoteInput) (*domain.Remote, error) {
	buttons := []*domain.Button{}
	for _, b := range input.Buttons {
		buttons = append(buttons, domain.Factory(b.Name, b.Tag))
	}
	return i.repo.CreateRemote(
		ctx,
		domain.RemoteFactory(input.Name, input.DeviveID, input.Tag, buttons),
	)
}

func (i *Interactor) GetRemote(ctx context.Context, input bdy.GetRemoteInput) (*domain.Remote, error) {
	return i.repo.ReadRemote(ctx, input.RemoteID)
}

func (i *Interactor) ListRemotes(ctx context.Context) ([]*domain.Remote, error) {
	return i.repo.ReadRemotes(ctx)
}

func (i *Interactor) UpdateRemote(ctx context.Context, input bdy.UpdateRemoteInput) error {
	r, err := i.repo.ReadRemote(ctx, input.RemoteID)
	if err != nil {
		return err
	}
	r.UpdateRemote(input.RemoteName, input.DeviceID)
	return i.repo.UpdateRemote(ctx, r)
}

func (i *Interactor) DeleteRemote(ctx context.Context, input bdy.DeleteRemoteInput) error {
	return i.repo.DeleteRemote(ctx, input.RemoteID)
}

func (i *Interactor) LearnIR(ctx context.Context, input bdy.LearnIRInput) error {
	b, err := i.repo.ReadButton(ctx, input.ButtonID)
	if err != nil {
		return err
	}
	b.LearnIR(input.IRData)
	return i.repo.UpdateButton(ctx, b)
}

func (i *Interactor) GetDevice(ctx context.Context, input bdy.GetDeivceInput) (*domain.Device, error) {
	return i.device.ReadDevice(ctx, input.DeviceID)
}

func (i *Interactor) ListDevices(ctx context.Context) ([]*domain.Device, error) {
	return i.device.ReadDevices(ctx)
}

func (i *Interactor) SendIR(ctx context.Context, input bdy.SendIRInput) error {
	return i.device.SendIR(ctx, input.ID, input.IRData)
}

func (i *Interactor) ReceiveIR(ctx context.Context, input bdy.ReceiveIRInput) (domain.IRData, error) {
	return i.device.ReceiveIR(ctx, input.ID)
}
