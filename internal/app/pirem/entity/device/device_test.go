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

func (m *Mock) SendIR(ctx context.Context, irdata ir.Data) error {
	time.Sleep(5 * time.Second)
	fmt.Println("sended")
	return nil
}

func (m *Mock) ReceiveIR(ctx context.Context) (ir.Data, error) {
	time.Sleep(5 * time.Second)
	irdata := &ir.RawData{}
	return irdata, nil
}

func TestNew(t *testing.T) {
	t.Run("device entitiy test", func(t *testing.T) {
		var wg sync.WaitGroup
		dev, err := New("1", "mock", Info{CanSend: true, CanReceive: true, DriverVersion: "0.0.1", FirmwareVersion: "0.0.1"}, &Mock{})
		if err != nil {
			return
		}

		sendIR := func(ctx context.Context, w *sync.WaitGroup) {
			defer w.Done()
			fmt.Println(dev.ReceiveIR(ctx))
		}

		ctx, c := context.WithTimeout(context.Background(), time.Second*2)

		wg.Add(1)
		go sendIR(ctx, &wg)
		wg.Add(1)
		go sendIR(ctx, &wg)
		wg.Wait()
		c()
	})

}
