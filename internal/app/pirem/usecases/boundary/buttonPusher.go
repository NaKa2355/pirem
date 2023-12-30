package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type PushButtonInput struct {
	ButtonId domain.ButtonID
}

type ButtonPusher interface {
	PushRemote(context.Context, PushButtonInput) error
}
