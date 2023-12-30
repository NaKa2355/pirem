package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type GetRemoteInput struct {
	RemoteID domain.RemoteID
}

type RemoteGetter interface {
	GetRemote(context.Context, GetRemoteInput) (*domain.Remote, error)
}
