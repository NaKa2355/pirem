package repository

import (
	"errors"
	"sync"

	entdev "github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	repo "github.com/NaKa2355/pirem/internal/app/pirem/usecases/repository"
)

type Repository struct {
	Devices map[string]*entdev.Device
	mu      sync.RWMutex
}

func New() *Repository {
	return &Repository{
		Devices: make(map[string]*entdev.Device),
		mu:      sync.RWMutex{},
	}
}

func (r *Repository) CreateDevice(dev *entdev.Device) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.Devices[dev.ID]; !ok {
		return repo.WrapErr(
			repo.CodeAlreadyExists,
			errors.New("device is already exists"),
		)
	}
	r.Devices[dev.ID] = dev
	return nil
}

func (r *Repository) ReadDevice(id string) (*entdev.Device, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	dev, ok := r.Devices[id]
	if !ok {
		return dev, repo.WrapErr(
			repo.CodeNotExist,
			errors.New("device is not exists"),
		)
	}
	return dev, nil
}

func (r *Repository) ReadDevices() ([]*entdev.Device, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	devs := make([]*entdev.Device, 0, len(r.Devices))
	for _, d := range r.Devices {
		devs = append(devs, d)
	}
	return devs, nil
}

func (r *Repository) DeleteDevice(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	dev, ok := r.Devices[id]
	if !ok {
		return repo.WrapErr(
			repo.CodeNotExist,
			errors.New("device is not exists"),
		)
	}

	delete(r.Devices, dev.ID)
	return nil
}
