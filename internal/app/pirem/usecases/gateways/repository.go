package gateway

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
)

type Repository interface {
	CreateRemote(ctx context.Context, a *remote.Remote) (*remote.Remote, error)

	ReadRemote(context.Context, remote.ID) (*remote.Remote, error)
	ReadRemotes(context.Context) ([]*remote.Remote, error)
	ReadButtons(ctx context.Context, appID remote.ID) ([]*button.Button, error)
	ReadButton(context.Context, button.ID) (*button.Button, error)

	UpdateRemote(context.Context, *remote.Remote) error
	UpdateButton(context.Context, *button.Button) error

	DeleteRemote(context.Context, remote.ID) error
}
