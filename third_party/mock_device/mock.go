package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NaKa2355/pirem/pkg/plugin/driver"
)

type MockDev struct{}

func GetDriver(conf json.RawMessage) (driver.Driver, error) {
	return &MockDev{}, nil
}

func (m *MockDev) SendIR(ctx context.Context, irdata driver.IRData) error {
	fmt.Println("send ir")
	return nil
}

func (m *MockDev) ReceiveIR(ctx context.Context) (driver.IRData, error) {
	fmt.Println("receive ir")
	return driver.IRData{CarrierFreqKiloHz: 40, PluseNanoSec: []uint32{10, 20, 30}}, nil
}

func (m *MockDev) GetInfo(ctx context.Context) (driver.DeviceInfo, error) {
	return driver.DeviceInfo{
		CanSend:         true,
		CanReceive:      true,
		DriverVersion:   "0.1",
		FirmwareVersion: "0.2",
	}, nil
}
