package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type SendIRInput struct {
	ID     domain.DeviceID
	IRData domain.IRData
}

type IRSender interface {
	SendIR(context.Context, SendIRInput) error
}
