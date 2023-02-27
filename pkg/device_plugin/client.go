package device_plugin

import (
	"context"
	"encoding/json"

	irdatav1 "github.com/NaKa2355/irdeck-proto/gen/go/common/irdata/v1"
	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
	pluginv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/plugin/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/empty"
)

type GRPCClient struct {
	client pluginv1.DevicePluginServiceClient
}

func (m *GRPCClient) SendRawIr(ctx context.Context, irData *irdatav1.RawIrData) error {
	_, err := m.client.SendRawIr(ctx, irData)
	return err
}

func (m *GRPCClient) ReceiveRawIr(ctx context.Context) (*irdatav1.RawIrData, error) {
	return m.client.ReceiveRawIr(ctx, &empty.Empty{})
}

func (m *GRPCClient) GetDeviceInfo(ctx context.Context) (*apiremv1.DeviceInfo, error) {
	return m.client.GetDeviceInfo(ctx, &empty.Empty{})
}

func (m *GRPCClient) GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error) {
	return m.client.GetDeviceStatus(ctx, &empty.Empty{})
}

func (m *GRPCClient) IsBusy(ctx context.Context) (bool, error) {
	resp, err := m.client.IsBusy(ctx, &emptypb.Empty{})
	return resp.IsBusy, err
}

func (m *GRPCClient) Init(ctx context.Context, conf json.RawMessage) error {
	_, err := m.client.Init(ctx, &pluginv1.DeviceConfig{JsonConfig: string(conf)})
	return err
}
