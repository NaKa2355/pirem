package plugin

import (
	"context"
	"encoding/json"

	apiremv1 "github.com/NaKa2355/pirem/gen/apirem/v1"
	pluginv1 "github.com/NaKa2355/pirem/gen/plugin/v1"

	"github.com/golang/protobuf/ptypes/empty"
)

type GRPCClient struct {
	client pluginv1.DevicePluginServiceClient
}

func (m *GRPCClient) SendRawIr(ctx context.Context, irData *apiremv1.RawIrData) error {
	_, err := m.client.SendRawIr(ctx, irData)
	return err
}

func (m *GRPCClient) ReceiveRawIr(ctx context.Context) (*apiremv1.RawIrData, error) {
	return m.client.ReceiveRawIr(ctx, &empty.Empty{})
}

func (m *GRPCClient) GetDeviceInfo(ctx context.Context) (*apiremv1.DeviceInfo, error) {
	return m.client.GetDeviceInfo(ctx, &empty.Empty{})
}

func (m *GRPCClient) GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error) {
	return m.client.GetDeviceStatus(ctx, &empty.Empty{})
}

func (m *GRPCClient) Init(ctx context.Context, conf json.RawMessage) error {
	_, err := m.client.Init(ctx, &pluginv1.DeviceConfig{JsonConfig: string(conf)})
	return err
}

func (m *GRPCClient) Drop() error {
	return nil
}
