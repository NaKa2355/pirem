package gateways

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type Repository interface {
	CreateRemote(ctx context.Context, a *domain.Remote) (*domain.Remote, error)

	ReadRemote(context.Context, domain.RemoteID) (*domain.Remote, error)
	ReadRemotes(context.Context) ([]*domain.Remote, error)
	ReadButtons(ctx context.Context, appID domain.RemoteID) ([]*domain.Button, error)
	ReadButton(context.Context, domain.ButtonID) (*domain.Button, error)

	ReadIRDataAndDeviceID(context.Context, domain.ButtonID) (domain.IRData, domain.DeviceID, error)
	UpdateRemote(context.Context, *domain.Remote) error
	LearnIR(context.Context, domain.ButtonID, domain.IRData) error
	DeleteRemote(context.Context, domain.RemoteID) error
}
