package plugin

import (
	"context"
	"encoding/json"

	apiremv1 "github.com/NaKa2355/pirem/pkg/apirem.v1"

	"github.com/golang/protobuf/ptypes/empty"
)

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl DeviceController
	apiremv1.UnimplementedDevicePluginServiceServer
}

func (m *GRPCServer) SendRawIr(ctx context.Context, irData *apiremv1.RawIrData) (*empty.Empty, error) {
	err := m.Impl.SendRawIr(ctx, irData)
	return &empty.Empty{}, err
}

func (m *GRPCServer) ReceiveRawIr(ctx context.Context, e *empty.Empty) (*apiremv1.RawIrData, error) {
	return m.Impl.ReceiveRawIr(ctx)
}

func (m *GRPCServer) GetDeviceInfo(ctx context.Context, e *empty.Empty) (*apiremv1.DeviceInfo, error) {
	return m.Impl.GetDeviceInfo(ctx)
}

func (m *GRPCServer) GetDeviceStatus(ctx context.Context, e *empty.Empty) (*apiremv1.DeviceStatus, error) {
	return m.Impl.GetDeviceStatus(ctx)
}

func (m *GRPCServer) Drop() error {
	return nil
}

func (m *GRPCServer) Init(ctx context.Context, config *apiremv1.DeviceConfig) (*empty.Empty, error) {
	err := m.Impl.Init(ctx, json.RawMessage(config.JsonConfig))
	return &empty.Empty{}, err
}
