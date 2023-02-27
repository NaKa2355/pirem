package device_plugin

import (
	"context"
	"fmt"

	pluginv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/plugin/v1"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var ErrPluginNotSupported error = fmt.Errorf("plugin not supported")

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "DEVICE_PLUGIN",
	MagicCookieValue: "device_controller",
}

var PluginMap = map[string]plugin.Plugin{
	"device_controller": &DevicePlugin{},
}

type DevicePlugin struct {
	plugin.Plugin
	Impl DeviceController
}

func (p *DevicePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pluginv1.RegisterDevicePluginServiceServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *DevicePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: pluginv1.NewDevicePluginServiceClient(c)}, nil
}
