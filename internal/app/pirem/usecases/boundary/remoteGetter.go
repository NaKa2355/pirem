package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
)

type GetRemoteInput struct {
	RemoteID remote.ID
}

type RemoteGetter interface {
	GetRemote(context.Context, GetRemoteInput) (*remote.Remote, error)
}
