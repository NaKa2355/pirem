package remote

import (
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
)

type Remote struct {
	ID       ID
	Name     Name
	DeviceID device.ID
	Tag      Tag
	Buttons  []*button.Button
}

func NewRemote(id ID, name Name, deviceID device.ID, tag Tag, buttons []*button.Button) *Remote {
	return &Remote{
		ID:       id,
		Name:     name,
		Tag:      tag,
		DeviceID: deviceID,
		Buttons:  buttons,
	}
}

func RemoteFactory(name Name, deviceID device.ID, tag Tag, buttons []*button.Button) *Remote {
	return NewRemote(NewID(), name, deviceID, tag, buttons)
}

func RestoreRemote(id ID, name Name, deviceID device.ID, tag Tag, buttons []*button.Button) *Remote {
	return &Remote{
		ID:       id,
		Name:     name,
		Tag:      tag,
		DeviceID: deviceID,
		Buttons:  buttons,
	}
}

func (r *Remote) UpdateRemote(name Name, deviceID device.ID) {
	r.Name = name
	r.DeviceID = deviceID
}
