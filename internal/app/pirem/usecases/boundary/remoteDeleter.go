package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type DeleteRemoteInput struct {
	RemoteID domain.RemoteID
}

type RemoteDeleter interface {
	DeleteRemote(context.Context, DeleteRemoteInput) error
}
