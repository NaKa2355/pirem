package main

import (
	"encoding/json"

	plugin "github.com/NaKa2355/pirem/pkg/plugin/driver"
	"github.com/NaKa2355/pirem/third_party/airer/internal/app/airer/device"
)

// build with this command
// $ go build -buildmode=plugin

//change plugin name below command.
//if you don't change, daemon might not able to load your plugin.
// $ cd ..; mv plugin new_name

func GetDriver(jsonConf json.RawMessage) (plugin.Driver, error) {
	return device.NewDevice(jsonConf)
}

func main() {}
