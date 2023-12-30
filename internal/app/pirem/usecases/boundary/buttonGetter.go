package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type GetButtonInput struct {
	ButtonID domain.ButtonID
}

type ButtonGetter interface {
	GetButton(context.Context, GetButtonInput) (*domain.Button, error)
}
