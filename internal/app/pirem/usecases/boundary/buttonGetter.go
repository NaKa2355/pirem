package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
)

type GetButtonInput struct {
	ButtonID button.ID
}

type ButtonGetter interface {
	GetButton(context.Context, GetButtonInput) (*button.Button, error)
}
