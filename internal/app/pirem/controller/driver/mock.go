package driver

import (
	"context"
	"fmt"

	"github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/ir"
)

var _ device.Driver = &Mock{}

type Mock struct {
	Info       device.Info
	sendErr    error
	receiveErr error
}

func NewMock(sendErr error, receiveErr error) *Mock {
	return &Mock{
		Info: device.Info{
			CanSend:         true,
			CanReceive:      true,
			DriverVersion:   "0.1",
			FirmwareVersion: "0.1",
		},
	}
}

func (m *Mock) GetDeviceInfo() *device.Info {
	return &m.Info
}

func (m *Mock) SendIR(ctx context.Context, irData ir.Data) error {
	fmt.Println("send ir")
	fmt.Println(irData)
	return convertErr(m.sendErr)
}

func (m *Mock) ReceiveIR(ctx context.Context) (ir.Data, error) {
	rawIRData := &ir.RawData{}

	rawIRData.CarrierFreqKiloHz = 40
	rawIRData.PluseNanoSec = []uint32{10, 20, 20, 30, 60}
	fmt.Println("receive ir")
	return rawIRData, convertErr(m.receiveErr)
}
