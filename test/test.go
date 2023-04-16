package main

import (
	"context"
	"encoding/json"
	"fmt"

	devctr "github.com/NaKa2355/pirem/internal/app/pirem/controller/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/repository"
	entdev "github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	interactor "github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
)

func main() {
	var err error = nil
	repo := repository.New()
	interactor := interactor.New(repo)

	driver, err := devctr.New("/home/rasp/programs/pirem/third_party/mock_device/mock.so", json.RawMessage{})
	if err != nil {
		fmt.Println(err)
		return
	}

	dev, err := entdev.New("1", "mock", driver.Info, driver)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = repo.CreateDevice(dev)
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := interactor.GetDeviceInfo(context.Background(), boundary.GetDeviceInput{ID: "1"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
}
