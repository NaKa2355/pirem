package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
)

type CreateRemoteInput struct {
	Name     remote.Name
	Tag      remote.Tag
	DeviveID device.ID
	Buttons  []struct {
		Name button.Name
		Tag  button.Tag
	}
}

type RemoteCreator interface {
	CreateRemote(context.Context, CreateRemoteInput) (*remote.Remote, error)
}
