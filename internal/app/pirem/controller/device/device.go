package device

import (
	"context"
	"encoding/json"
	"os/exec"

	apiremv1 "github.com/NaKa2355/pirem/gen/apirem/v1"
	dev_plugin "github.com/NaKa2355/pirem/pkg/device_plugin"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

var _ dev_plugin.DeviceController = &Device{}

type Device struct {
	devCtrl dev_plugin.DeviceController
	client  *plugin.Client
}

func New(pluginPath string, conf json.RawMessage, logger hclog.Logger) (*Device, error) {
	dev := &Device{}
	client := plugin.NewClient(
		&plugin.ClientConfig{
			HandshakeConfig: dev_plugin.Handshake,
			Plugins:         dev_plugin.PluginMap,
			Cmd:             exec.Command(pluginPath),
			Logger:          logger,
			AllowedProtocols: []plugin.Protocol{
				plugin.ProtocolGRPC,
			},
		},
	)

	rpcClient, err := client.Client()
	if err != nil {
		return dev, err
	}

	raw, err := rpcClient.Dispense("device_controller")
	if err != nil {
		client.Kill()
		return dev, err
	}

	devCtrl, ok := raw.(dev_plugin.DeviceController)
	if !ok {
		client.Kill()
		return dev, dev_plugin.ErrPluginNotSupported
	}

	if err := devCtrl.Init(context.Background(), conf); err != nil {
		client.Kill()
		return dev, err
	}

	dev.devCtrl = devCtrl
	dev.client = client
	return dev, nil
}

func (d *Device) SendRawIr(ctx context.Context, irData *apiremv1.RawIrData) error {
	return d.devCtrl.SendRawIr(ctx, irData)
}

func (d *Device) ReceiveRawIr(ctx context.Context) (*apiremv1.RawIrData, error) {
	return d.devCtrl.ReceiveRawIr(ctx)
}

func (d *Device) GetDeviceInfo(ctx context.Context) (*apiremv1.DeviceInfo, error) {
	return d.devCtrl.GetDeviceInfo(ctx)
}

func (d *Device) GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error) {
	return d.devCtrl.GetDeviceStatus(ctx)
}

func (d *Device) IsBusy(ctx context.Context) (bool, error) {
	return d.devCtrl.IsBusy(ctx)
}

func (d *Device) Init(ctx context.Context, conf json.RawMessage) error {
	return d.devCtrl.Init(ctx, conf)
}

func (d *Device) Drop() error {
	d.client.Kill()
	return nil
}
