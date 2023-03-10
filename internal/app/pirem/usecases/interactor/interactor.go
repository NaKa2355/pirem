package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"sync"

	irdatav1 "github.com/NaKa2355/irdeck-proto/gen/go/common/irdata/v1"
	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/device"
	entity_dev "github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	dev_usecases "github.com/NaKa2355/pirem/internal/app/pirem/usecases/device"
	pirem_err "github.com/NaKa2355/pirem/pkg/error"
	"github.com/hashicorp/go-hclog"
)

const MsgDevNotFound = "device(id=%s) is not exist: %w"
const MsgDevNotSupported = "device(id*%s) does not support this command"
const IDRegExp = "^[*-~]*$"

type Interactor struct {
	mu      sync.RWMutex
	devices map[string]dev_usecases.DeviceController
	logger  hclog.Logger
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

func New(logger hclog.Logger) *Interactor {
	e := Interactor{}
	e.devices = make(map[string]dev_usecases.DeviceController)
	e.logger = logger
	return &e
}

// add device
func (e *Interactor) AddDevice(id string, name string, pluginFilePath string, conf json.RawMessage) error {
	if err := validateDeviceID(IDRegExp, id); err != nil {
		return err
	}

	devCtrl, err := device.New(pluginFilePath, conf, e.logger)
	if err != nil {
		return err
	}

	dev, err := entity_dev.New(id, name, devCtrl)
	if err != nil {
		devCtrl.Drop()
		return err
	}

	e.mu.Lock()
	e.devices[id] = dev
	e.mu.Unlock()
	return nil
}

// get all devices information
func (e *Interactor) GetAllDeviceInfo(ctx context.Context) ([]*apiremv1.DeviceInfo, error) {
	var err error = nil
	var infoList = make([]*apiremv1.DeviceInfo, 0, 2)
	var info *apiremv1.DeviceInfo

	e.mu.RLock()
	defer e.mu.RUnlock()

	for _, device := range e.devices {
		//entity???GetDeviceInfo???error?????????nil??????????????????????????????????????????
		info, _ = device.GetDeviceInfo(ctx)
		infoList = append(infoList, info)
	}

	return infoList, err
}

// get device information
func (e *Interactor) GetDeviceInfo(ctx context.Context, id string) (*apiremv1.DeviceInfo, error) {
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

	return device.GetDeviceInfo(ctx)
}

func (e *Interactor) GetDeviceStatus(ctx context.Context, id string) (*apiremv1.DeviceStatus, error) {
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

func (e *Interactor) IsBusy(ctx context.Context, id string) (bool, error) {
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

func (e *Interactor) SendRawIr(ctx context.Context, id string, ir_data *irdatav1.RawIrData) error {
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

	err = device.SendRawIr(ctx, ir_data)
	if err != nil {
		err = fmt.Errorf("%s: %w", err, pirem_err.ErrDeviceInternal)
	}
	return err
}

func (e *Interactor) ReceiveRawIr(ctx context.Context, id string) (*irdatav1.RawIrData, error) {
	var err error = nil
	var irData *irdatav1.RawIrData

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

	irData, err = device.ReceiveRawIr(ctx)
	if err != nil {
		err = fmt.Errorf("%s: %w", err, pirem_err.ErrDeviceInternal)
	}
	return irData, err
}

// free resources
func (e *Interactor) Drop() {
	e.mu.RLock()
	for deviceId, device := range e.devices {
		device.Drop()
		delete(e.devices, deviceId)
	}
	e.mu.RUnlock()
}
