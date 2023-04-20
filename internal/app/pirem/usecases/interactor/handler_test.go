package interactor

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	driver "github.com/NaKa2355/pirem/internal/app/pirem/controller/driver"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/repository"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
)

func TestNew(t *testing.T) {
	t.Run("test", func(tt *testing.T) {
		fmt.Println("hello")
		repo := repository.New()
		driv, err := driver.New("/home/rasp/programs/pirem/third_party/mock_device/mock.so", json.RawMessage{})
		if err != nil {
			fmt.Println(err)
			return
		}
		dev, err := device.New("1", "test", device.Info{CanSend: true, CanReceive: true, DriverVersion: "0.1.0", FirmwareVersion: "0.0.1"}, driv)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(repo.CreateDevice(dev))
		in := New(repo)
		fmt.Println(in.SendIR(context.Background(), boundary.SendIRInput{ID: "1", IRData: boundary.RawIRData{}}))
		fmt.Println(in.ReceiveIR(context.Background(), boundary.ReceiveIRInput{ID: "1"}))
	})
}
