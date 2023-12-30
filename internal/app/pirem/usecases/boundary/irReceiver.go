package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type ReceiveIRInput struct {
	ID domain.DeviceID
}

type IRReceiver interface {
	ReceiveIR(context.Context, ReceiveIRInput) (domain.IRData, error)
}
