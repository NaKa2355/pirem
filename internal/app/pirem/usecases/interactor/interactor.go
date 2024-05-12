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
	var button *domain.Button
	err := i.repo.Transaction(ctx, func(ctx context.Context, repo gateway.Repository) (err error) {
		button, err = repo.ReadButton(ctx, input.ButtonID)
		return err
	})
	return button, err
}

func (i *Interactor) PushRemote(ctx context.Context, input bdy.PushButtonInput) error {
	var irData domain.IRData
	var deviceID domain.DeviceID

	err := i.repo.Transaction(ctx, func(ctx context.Context, repo gateway.Repository) (err error) {
		irData, deviceID, err = repo.ReadIRDataAndDeviceID(ctx, input.ButtonId)
		return err
	})

	if err != nil {
		return err
	}
	return i.device.SendIR(ctx, deviceID, irData)
}

func (i *Interactor) CreateRemote(ctx context.Context, input bdy.CreateRemoteInput) (*domain.Remote, error) {
	buttons := []*domain.Button{}
	for _, b := range input.Buttons {
		buttons = append(buttons, domain.ButtonFactory(b.Name, b.Tag, input.DeviveID))
	}
	var remote *domain.Remote
	err := i.repo.Transaction(ctx, func(ctx context.Context, repo gateway.Repository) (err error) {
		remote, err = repo.CreateRemote(ctx, domain.RemoteFactory(input.Name, input.DeviveID, input.Tag, buttons))
		return err
	})
	return remote, err
}

func (i *Interactor) GetRemote(ctx context.Context, input bdy.GetRemoteInput) (*domain.Remote, error) {
	var remote *domain.Remote
	err := i.repo.Transaction(ctx, func(ctx context.Context, repo gateway.Repository) (err error) {
		remote, err = repo.ReadRemote(ctx, input.RemoteID)
		return err
	})
	return remote, err
}

func (i *Interactor) ListRemotes(ctx context.Context) ([]*domain.Remote, error) {
	var remotes []*domain.Remote
	err := i.repo.Transaction(ctx, func(ctx context.Context, repo gateway.Repository) (err error) {
		remotes, err = repo.ReadRemotes(ctx)
		return err
	})
	return remotes, err
}

func (i *Interactor) UpdateRemote(ctx context.Context, input bdy.UpdateRemoteInput) error {
	return i.repo.Transaction(ctx, func(ctx context.Context, repo gateway.Repository) error {
		r, err := repo.ReadRemote(ctx, input.RemoteID)
		if err != nil {
			return err
		}
		r.UpdateRemote(input.RemoteName, input.DeviceID)
		return repo.UpdateRemote(ctx, r)
	})
}

func (i *Interactor) DeleteRemote(ctx context.Context, input bdy.DeleteRemoteInput) error {
	return i.repo.Transaction(ctx, func(ctx context.Context, repo gateway.Repository) error {
		return repo.DeleteRemote(ctx, input.RemoteID)
	})
}

func (i *Interactor) LearnIR(ctx context.Context, input bdy.LearnIRInput) error {
	return i.repo.Transaction(ctx, func(ctx context.Context, repo gateway.Repository) error {
		return repo.LearnIR(ctx, input.ButtonID, input.IRData)
	})
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

func (i *Interactor) GetIR(ctx context.Context, in bdy.GetIRInput) (bdy.GetIROutput, error) {
	irdata, deviceID, err := i.repo.ReadIRDataAndDeviceID(ctx, in.ButtonID)
	return bdy.GetIROutput{
		IRData:   irdata,
		DeviceID: deviceID,
	}, err
}
