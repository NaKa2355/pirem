package interactor

import (
	"context"

	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
)

func New(repo Repository) *Interactor {
	i := &Interactor{}
	i.devsRepo = repo
	return i
}

func (i *Interactor) GetDevicesInfo(ctx context.Context) (out bdy.GetDevicesInfoOutput, err error) {
	return i.getDevicesInfo(ctx)
}

func (i *Interactor) GetDeviceInfo(ctx context.Context, in bdy.GetDeviceInput) (out bdy.GetDeviceInfoOutput, err error) {
	return i.getDeviceInfo(ctx, in)
}

func (i *Interactor) SendIR(ctx context.Context, in bdy.SendIRInput) (err error) {
	return i.sendIR(ctx, in)
}

func (i *Interactor) ReceiveIR(ctx context.Context, in bdy.ReceiveIRInput) (out bdy.IRData, err error) {
	return i.receiveIR(ctx, in)
}
