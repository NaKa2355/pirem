package device

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/NaKa2355/pirem/internal/app/pirem/entity/ir"
)

type Mock struct{}

var _ Driver = &Mock{}

func (m *Mock) GetDeviceInfo() *Info {
	return &Info{CanSend: true, CanReceive: true, DriverVersion: "0.0.1", FirmwareVersion: "0.0.1"}
}

func (m *Mock) SendIR(ctx context.Context, irdata ir.Data) error {
	time.Sleep(1 * time.Second)
	fmt.Println("sended")
	return nil
}

func (m *Mock) ReceiveIR(ctx context.Context) (ir.Data, error) {
	time.Sleep(5 * time.Second)
	irdata := &ir.RawData{}
	return irdata, nil
}

func TestParallelDeviceReq(t *testing.T) {
	t.Run("device entitiy test", func(t *testing.T) {
		var wg sync.WaitGroup
		dev, err := New("1", "mock", &Mock{}, 5*time.Second)
		if err != nil {
			return
		}

		sendIr := func(ctx context.Context, w *sync.WaitGroup) {
			wg.Add(1)
			go func() {
				defer w.Done()
				fmt.Println(dev.SendIR(ctx, &ir.RawData{}))
			}()
		}

		receiveIR := func(ctx context.Context, w *sync.WaitGroup) {
			wg.Add(1)
			go func() {
				defer w.Done()
				fmt.Println(dev.ReceiveIR(ctx))
			}()
		}

		//ctx, c := context.WithTimeout(context.Background(), time.Second*2)
		sendIr(context.Background(), &wg)
		sendIr(context.Background(), &wg)
		receiveIR(context.Background(), &wg)
		receiveIR(context.Background(), &wg)
		wg.Wait()
		//c()
	})
}
