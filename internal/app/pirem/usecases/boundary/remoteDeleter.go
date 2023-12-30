package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
)

type DeleteRemoteInput struct {
	RemoteID remote.ID
}

type RemoteDeleter interface {
	DeleteRemote(context.Context, DeleteRemoteInput) error
}
