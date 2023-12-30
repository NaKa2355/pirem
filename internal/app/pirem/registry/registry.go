package registry

import (
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/gateways"
)

type DeviceFactory func() (controllers.IRDevice, error)
type RepositoryFactory func() (gateways.Repository, error)
type BoudaryFactory func(device controllers.IRDevice, repo gateways.Repository) (boundary.Boundary, error)
type EntryPoint func(boundary boundary.Boundary)

type Handlers interface {
	DeviceFactory() (controllers.IRDevice, error)
	RepositoryFactory() (gateways.Repository, error)
	BoudaryFactory(device controllers.IRDevice, repo gateways.Repository) (boundary.Boundary, error)
	EntryPoint(boundary boundary.Boundary)
}

type Registry struct {
	h Handlers
	b boundary.Boundary
}

func NewRegistry(h Handlers) *Registry {
	return &Registry{
		h: h,
		b: nil,
	}
}

func (r *Registry) SolveDependencies() error {
	c, err := r.h.DeviceFactory()
	if err != nil {
		return err
	}
	repo, err := r.h.RepositoryFactory()
	if err != nil {
		return err
	}
	b, err := r.h.BoudaryFactory(c, repo)
	if err != nil {
		return err
	}
	r.b = b
	return nil
}

func (r *Registry) Start() {
	if r.b == nil {
		return
	}
	r.h.EntryPoint(r.b)
}
