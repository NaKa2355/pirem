package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/irdata"
)

type ReceiveIRInput struct {
	ID device.ID
}

type IRReceiver interface {
	ReceiveIR(context.Context, ReceiveIRInput) (irdata.IRData, error)
}
