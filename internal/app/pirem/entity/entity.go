package entity

import (
	"context"
	"fmt"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	apiremv1 "github.com/NaKa2355/pirem/pkg/apirem.v1"
	pirem_err "github.com/NaKa2355/pirem/pkg/error"
	"regexp"
	"sync"
)

const MsgDevNotFound = "device(id=%s) is not exist: %w"
const MsgDevNotSupported = "device(id*%s) does not support this command"
const IDRegExp = "^[*-~]*$"

type Entity struct {
	mu      sync.RWMutex
	devices map[string]usecases.DeviceController
	apiremv1.UnimplementedPiRemServiceServer

	usecases.EntityController
}

// validate device id
func validateDeviceID(pattern string, deviceID string) error {
	match, err := regexp.MatchString(pattern, deviceID)
	if err != nil {
		err = fmt.Errorf("faild to validate: %w", err)
		return err
	}
	if !match {
		err = fmt.Errorf("invaild input (id=%s) you must follow this format (%s)", deviceID, pattern)
		err = fmt.Errorf("%s: %w", err, pirem_err.ErrInvaildArgument)
		return err
	}
	return nil
}

func New() *Entity {
	e := Entity{}
	e.devices = make(map[string]usecases.DeviceController)
	return &e
}

// add device
func (e *Entity) AddDevice(dev usecases.DeviceController) error {
	info := dev.GetDeviceInfo(context.Background())
	id := info.GetId()

	if err := validateDeviceID(IDRegExp, id); err != nil {
		return err
	}

	e.mu.Lock()
	e.devices[id] = dev
	e.mu.Unlock()

	return nil
}

// get all devices information
func (e *Entity) GetAllDeviceInfo(ctx context.Context) ([]*apiremv1.DeviceInfo, error) {
	var err error = nil
	var deviceInfo = make([]*apiremv1.DeviceInfo, 0, 2)

	e.mu.RLock()
	for _, device := range e.devices {
		deviceInfo = append(deviceInfo, device.GetDeviceInfo(ctx))
	}
	e.mu.RUnlock()

	return deviceInfo, err
}

// get device information
func (e *Entity) GetDeviceInfo(ctx context.Context, id string) (*apiremv1.DeviceInfo, error) {
	var err error = nil
	var info *apiremv1.DeviceInfo

	if err := validateDeviceID(IDRegExp, id); err != nil {
		return info, err
	}

	e.mu.RLock()
	device, ok := e.devices[id]
	e.mu.RUnlock()

	if !ok {
		err = fmt.Errorf(MsgDevNotFound, id, pirem_err.ErrDeviceNotFound)
		return info, err
	}

	return device.GetDeviceInfo(ctx), err
}

func (e *Entity) GetDeviceStatus(ctx context.Context, id string) (*apiremv1.DeviceStatus, error) {
	var err error = nil
	var status *apiremv1.DeviceStatus

	if err := validateDeviceID(IDRegExp, id); err != nil {
		return status, err
	}

	e.mu.RLock()
	device, ok := e.devices[id]
	e.mu.RUnlock()

	if !ok {
		err = fmt.Errorf(MsgDevNotFound, id, pirem_err.ErrDeviceNotFound)
		return status, err
	}

	status, err = device.GetDeviceStatus(ctx)
	if err != nil {
		err = fmt.Errorf("%s: %w", err, pirem_err.ErrDeviceInternal)
	}
	return status, err
}

func (e *Entity) IsBusy(ctx context.Context, id string) (bool, error) {
	var err error = nil
	var isBusy bool

	if err := validateDeviceID(IDRegExp, id); err != nil {
		return isBusy, err
	}

	e.mu.RLock()
	device, ok := e.devices[id]
	e.mu.RUnlock()

	if !ok {
		err = fmt.Errorf(MsgDevNotFound, id, pirem_err.ErrDeviceNotFound)
		return isBusy, err
	}

	isBusy, err = device.IsBusy(ctx)
	if err != nil {
		err = fmt.Errorf("%s: %w", err, pirem_err.ErrDeviceInternal)
	}
	return isBusy, err
}

func (e *Entity) SendRawIr(ctx context.Context, id string, ir_data *apiremv1.RawIrData) error {
	var err error = nil

	if err := validateDeviceID(IDRegExp, id); err != nil {
		return err
	}

	e.mu.RLock()
	device, ok := e.devices[id]
	e.mu.RUnlock()

	if !ok {
		err = fmt.Errorf(MsgDevNotFound, id, pirem_err.ErrDeviceNotFound)
		return err
	}

	err = device.SendIr(ctx, ir_data)
	if err != nil {
		err = fmt.Errorf("%s: %w", err, pirem_err.ErrDeviceInternal)
	}
	return err
}

func (e *Entity) ReceiveRawIr(ctx context.Context, id string) (*apiremv1.RawIrData, error) {
	var err error = nil
	var irData *apiremv1.RawIrData

	if err := validateDeviceID(IDRegExp, id); err != nil {
		return irData, err
	}

	e.mu.RLock()
	device, ok := e.devices[id]
	e.mu.RUnlock()

	if !ok {
		err = fmt.Errorf(MsgDevNotFound, id, pirem_err.ErrDeviceNotFound)
		return irData, err
	}

	irData, err = device.ReceiveIr(ctx)
	if err != nil {
		err = fmt.Errorf("%s: %w", err, pirem_err.ErrDeviceInternal)
	}
	return irData, err
}

// free resources
func (e *Entity) Drop() {
	e.mu.RLock()
	for deviceId, device := range e.devices {
		device.Drop()
		delete(e.devices, deviceId)
	}
	e.mu.RUnlock()
}
