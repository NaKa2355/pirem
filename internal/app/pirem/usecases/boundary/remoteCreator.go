package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type CreateRemoteInput struct {
	Name     domain.RemoteName
	Tag      domain.RemoteTag
	DeviveID domain.DeviceID
	Buttons  []struct {
		Name domain.ButtonName
		Tag  domain.ButtonTag
	}
}

type RemoteCreator interface {
	CreateRemote(context.Context, CreateRemoteInput) (*domain.Remote, error)
}
