package mockdevice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NaKa2355/pirem/pkg/module/v1"
)

type Plugin struct {
}

func (p *Plugin) NewDriver(conf json.RawMessage) (module.Driver, error) {
	return &MockDev{}, nil
}

type MockDev struct{}

func (m *MockDev) SendIR(ctx context.Context, irdata *module.IRData) error {
	fmt.Println("send ir")
	return nil
}

func (m *MockDev) ReceiveIR(ctx context.Context) (*module.IRData, error) {
	fmt.Println("receive ir")
	return &module.IRData{CarrierFreqKiloHz: 40, PluseNanoSec: []uint32{10, 20, 30}}, nil
}

func (m *MockDev) GetInfo(ctx context.Context) (*module.DeviceInfo, error) {
	return &module.DeviceInfo{
		CanSend:         true,
		CanReceive:      true,
		DriverVersion:   "0.1",
		FirmwareVersion: "0.2",
	}, nil
}

func (m *MockDev) Drop() error {
	return nil
}
