package interactor

import (
	"context"
	"time"

	"github.com/NaKa2355/pirem/internal/app/pirem/entity"
	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	repo "github.com/NaKa2355/pirem/internal/app/pirem/usecases/repository"
	"github.com/NaKa2355/pirem/pkg/module/v1"
)

func New(repo Repository, mutexLockDeadline time.Duration) *Interactor {
	i := &Interactor{}
	i.mutexLockDeadline = mutexLockDeadline
	i.devsRepo = repo
	return i
}

func convertErr(err error) error {
	if err == nil {
		return nil
	}

	var code bdy.ErrCode
	switch _err := err.(type) {
	case *repo.Error:
		switch _err.Code {
		case repo.CodeAlreadyExists:
			code = bdy.CodeAlreadyExists
		case repo.CodeNotExist:
			code = bdy.CodeNotExist
		}

	case *module.Error:
		switch _err.Code {
		case module.CodeBusy:
			code = bdy.CodeBusy
		case module.CodeInvaildInput:
			code = bdy.CodeInvaildInput
		case module.CodeTimeout:
			code = bdy.CodeTimeout
		}

	case *entity.Error:
		switch _err.Code {
		case entity.CodeNotSupported:
			code = bdy.CodeNotSupported
		case entity.CodeDeviceBusy:
			code = bdy.CodeBusy
		}
	}

	return bdy.WrapErr(code, err)
}

func (i *Interactor) GetDevicesInfo(ctx context.Context) (out bdy.GetDevicesInfoOutput, err error) {
	out, err = i.getDevicesInfo(ctx)
	return out, convertErr(err)
}

func (i *Interactor) GetDeviceInfo(ctx context.Context, in bdy.GetDeviceInput) (out bdy.GetDeviceInfoOutput, err error) {
	out, err = i.getDeviceInfo(ctx, in)
	return out, convertErr(err)
}

func (i *Interactor) SendIR(ctx context.Context, in bdy.SendIRInput) (err error) {
	err = i.sendIR(ctx, in)
	return convertErr(err)
}

func (i *Interactor) ReceiveIR(ctx context.Context, in bdy.ReceiveIRInput) (out bdy.IRData, err error) {
	out, err = i.receiveIR(ctx, in)
	return out, convertErr(err)
}

func (i *Interactor) AddDevice(ctx context.Context, in bdy.AddDeviceInput) (err error) {
	err = i.addDevice(ctx, in, i.mutexLockDeadline)
	return convertErr(err)
}
