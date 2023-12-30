package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
)

type PushButtonInput struct {
	ButtonId button.ID
}

type ButtonPusher interface {
	PushRemote(context.Context, PushButtonInput) error
}
