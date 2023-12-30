package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/irdata"
)

type SendIRInput struct {
	ID     device.ID
	IRData irdata.IRData
}

type IRSender interface {
	SendIR(context.Context, SendIRInput) error
}
