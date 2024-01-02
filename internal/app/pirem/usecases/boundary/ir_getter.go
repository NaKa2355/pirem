package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type GetIRInput struct {
	ButtonID domain.ButtonID
}

type GetIROutput struct {
	IRData   domain.IRData
	DeviceID domain.DeviceID
}

type IRGetter interface {
	GetIR(context.Context, GetIRInput) (GetIROutput, error)
}
