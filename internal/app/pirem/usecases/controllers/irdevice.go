package controllers

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type IRDevice interface {
	ReadDevices(ctx context.Context) ([]*domain.Device, error)
	ReadDevice(ctx context.Context, id domain.DeviceID) (*domain.Device, error)

	SendIR(ctx context.Context, id domain.DeviceID, data domain.IRData) error
	ReceiveIR(tx context.Context, id domain.DeviceID) (domain.IRData, error)
}
