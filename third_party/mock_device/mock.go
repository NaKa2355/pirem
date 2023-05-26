package mockdevice

import (
	"context"
	"encoding/json"
	"fmt"

	plugin "github.com/NaKa2355/pirem/pkg/plugin/v1"
)

type Plugin struct {
}

func (p *Plugin) NewDriver(conf json.RawMessage) (plugin.Driver, error) {
	return &MockDev{}, nil
}

type MockDev struct{}

func (m *MockDev) SendIR(ctx context.Context, irdata *plugin.IRData) error {
	fmt.Println("send ir")
	return nil
}

func (m *MockDev) ReceiveIR(ctx context.Context) (*plugin.IRData, error) {
	fmt.Println("receive ir")
	return &plugin.IRData{CarrierFreqKiloHz: 40, PluseNanoSec: []uint32{10, 20, 30}}, nil
}

func (m *MockDev) GetInfo(ctx context.Context) (*plugin.DeviceInfo, error) {
	return &plugin.DeviceInfo{
		CanSend:         true,
		CanReceive:      true,
		DriverVersion:   "0.1",
		FirmwareVersion: "0.2",
	}, nil
}

func (m *MockDev) Drop() error {
	return nil
}
