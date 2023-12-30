package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
)

type UpdateRemoteInput struct {
	RemoteID   remote.ID
	RemoteName remote.Name
	DeviceID   device.ID
}

type RemoteUpdater interface {
	UpdateRemote(context.Context, UpdateRemoteInput) error
}
