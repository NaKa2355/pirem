package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type UpdateRemoteInput struct {
	RemoteID   domain.RemoteID
	RemoteName domain.RemoteName
	DeviceID   domain.DeviceID
}

type RemoteUpdater interface {
	UpdateRemote(context.Context, UpdateRemoteInput) error
}
