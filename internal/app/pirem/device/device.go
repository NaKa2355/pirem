package device

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"

	apiremv1 "github.com/NaKa2355/pirem/gen/apirem/v1"
	"github.com/NaKa2355/pirem/pkg/device_plugin"
	"github.com/hashicorp/go-hclog"
	go_plugin "github.com/hashicorp/go-plugin"
)

var ErrPluginNotSupported error = fmt.Errorf("plugin not supported")

type Device struct {
	Information *apiremv1.DeviceInfo
	controller  plugin.DeviceController
	mu          sync.Mutex
	client      *go_plugin.Client
	usecases.DeviceController
}

// check the device supports sending infraread
func canSend(serviceType apiremv1.DeviceInfo_ServiceType) bool {
	if serviceType == apiremv1.DeviceInfo_SERVICE_TYPE_SEND || serviceType == apiremv1.DeviceInfo_SERVICE_TYPE_SEND_RECEIVE {
		return true
	}
	return false
}

// check the device supports receiving infraread
func canReceive(serviceType apiremv1.DeviceInfo_ServiceType) bool {
	if serviceType == apiremv1.DeviceInfo_SERVICE_TYPE_RECEIVE || serviceType == apiremv1.DeviceInfo_SERVICE_TYPE_SEND_RECEIVE {
		return true
	}
	return false
}

func New(id string, name string, dev_ctrler plugin.DeviceController, client *go_plugin.Client) (*Device, error) {
	dev := &Device{}
	var err error
	ctx := context.Background()
	dev.controller = dev_ctrler
	dev.Information, err = dev.controller.GetDeviceInfo(ctx)
	dev.Information.Id = id
	dev.client = client
	dev.Information.Name = name
	return dev, err
}

func NewFromPlugin(id string, name string, conf json.RawMessage, pluginPath string, logger hclog.Logger) (*Device, error) {
	dev := &Device{}
	client := go_plugin.NewClient(
		&go_plugin.ClientConfig{
			HandshakeConfig: plugin.Handshake,
			Plugins:         plugin.PluginMap,
			Cmd:             exec.Command(pluginPath),
			Logger:          logger,
			AllowedProtocols: []go_plugin.Protocol{
				go_plugin.ProtocolGRPC,
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

	devCtrl, ok := raw.(plugin.DeviceController)
	if !ok {
		client.Kill()
		return dev, plugin.ErrPluginNotSupported
	}

	if err := devCtrl.Init(context.Background(), conf); err != nil {
		client.Kill()
		return dev, err
	}
	return New(id, name, devCtrl, client)
}

func (d *Device) GetDeviceInfo(ctx context.Context) *apiremv1.DeviceInfo {
	return d.Information
}

func (d *Device) IsBusy(ctx context.Context) (bool, error) {
	isBusy := !d.mu.TryLock()
	if !isBusy {
		d.mu.Unlock()
	}
	return isBusy, nil
}

func (d *Device) GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error) {
	var devStatus *apiremv1.DeviceStatus
	d.mu.Lock()
	defer d.mu.Unlock()
	select {
	case <-ctx.Done():
		return devStatus, ctx.Err()
	default:
		return d.controller.GetDeviceStatus(ctx)
	}
}

func (d *Device) SendIr(ctx context.Context, ir_data *apiremv1.RawIrData) error {
	if !canSend(d.Information.Service) {
		return fmt.Errorf("this device does not support sending")
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		err := d.controller.SendRawIr(ctx, ir_data)
		if err != nil {
			return err
		}

		//interval time to avoid conflict of data
		time.Sleep(200 * time.Millisecond)
		return nil
	}
}

func (d *Device) ReceiveIr(ctx context.Context) (*apiremv1.RawIrData, error) {
	var ir_data *apiremv1.RawIrData
	if !canReceive(d.Information.Service) {
		return ir_data, fmt.Errorf("this device does not support receiving")
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	select {
	case <-ctx.Done():
		return ir_data, ctx.Err()
	default:
		return d.controller.ReceiveRawIr(ctx)
	}
}

func (d *Device) Drop() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.client.Kill()
	return nil
}
