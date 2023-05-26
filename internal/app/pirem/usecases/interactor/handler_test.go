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
	"github.com/NaKa2355/pirem/pkg/plugin/v1"
	mockdevice "github.com/NaKa2355/pirem/third_party/mock_device"
)

func TestNew(t *testing.T) {
	t.Run("test", func(tt *testing.T) {
		fmt.Println("hello")
		repo := repository.New()
		driv, err := driver.New("mock", json.RawMessage{}, map[string]plugin.Plugin{"mock": &mockdevice.Plugin{}})
		if err != nil {
			fmt.Println(err)
			return
		}
		dev, err := device.New("1", "test", driv)
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
