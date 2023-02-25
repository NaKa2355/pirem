package plugin

import (
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

var ErrPluginNotSupported error = fmt.Errorf("no supported plugin")

func LoadPlugin(pluginPath string) (DeviceController, *plugin.Client, error) {
	var err error = nil
	var devCtrl DeviceController

	client := plugin.NewClient(
		&plugin.ClientConfig{
			HandshakeConfig: Handshake,
			Plugins:         PluginMap,
			Cmd:             exec.Command(pluginPath),
			AllowedProtocols: []plugin.Protocol{
				plugin.ProtocolGRPC,
			},
		},
	)

	rpcClient, err := client.Client()
	if err != nil {
		return devCtrl, client, err
	}

	raw, err := rpcClient.Dispense("device_controller")
	if err != nil {
		client.Kill()
		return devCtrl, client, err
	}

	devCtrl, ok := raw.(DeviceController)
	if !ok {
		client.Kill()
		return devCtrl, client, ErrPluginNotSupported
	}

	return devCtrl, client, nil
}
