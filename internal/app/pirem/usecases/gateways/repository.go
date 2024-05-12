package gateways

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type TransactionContext struct {
	Transaction interface{}
}

type Repository interface {
	Transaction(ctx context.Context, f func(context.Context, Repository) error) error

	CreateRemote(ctx context.Context, a *domain.Remote) (*domain.Remote, error)

	ReadRemote(ctx context.Context, remoteID domain.RemoteID) (*domain.Remote, error)
	ReadRemotes(ctx context.Context) ([]*domain.Remote, error)
	ReadButtons(ctx context.Context, remoteID domain.RemoteID) ([]*domain.Button, error)
	ReadButton(ctx context.Context, buttonID domain.ButtonID) (*domain.Button, error)

	ReadIRDataAndDeviceID(ctx context.Context, buttonID domain.ButtonID) (domain.IRData, domain.DeviceID, error)
	UpdateRemote(ctx context.Context, remote *domain.Remote) error
	LearnIR(ctx context.Context, button domain.ButtonID, irData domain.IRData) error
	DeleteRemote(ctx context.Context, remoteID domain.RemoteID) error
}
