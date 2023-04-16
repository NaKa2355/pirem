package interactor

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/entity"
	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	repo "github.com/NaKa2355/pirem/internal/app/pirem/usecases/repository"
)

func New(repo Repository) *Interactor {
	i := &Interactor{}
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

	case *entity.Error:
		switch _err.Code {
		case entity.CodeBusy:
			code = bdy.CodeBusy
		case entity.CodeInternal:
			code = bdy.CodeInternal
		case entity.CodeInvaildInput:
			code = bdy.CodeInvaildInput
		case entity.CodeNotSupported:
			code = bdy.CodeNotSupported
		case entity.CodeTimeout:
			code = bdy.CodeTimeout
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
