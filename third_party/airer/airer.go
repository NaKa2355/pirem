package airer

import (
	"encoding/json"

	"github.com/NaKa2355/pirem/pkg/module/v1"
	"github.com/NaKa2355/pirem/third_party/airer/internal/app/airer/device"
)

type Module struct{}

func (m *Module) NewDriver(jsonConf json.RawMessage) (module.Driver, error) {
	return device.NewDevice(jsonConf)
}
