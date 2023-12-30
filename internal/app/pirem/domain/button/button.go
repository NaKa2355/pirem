package button

import (
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/irdata"
)

type Button struct {
	ID       ID
	Name     Name
	Tag      Tag
	IRData   irdata.IRData
	DeviceID device.ID
}

func New(id ID, name Name, tag Tag, irData irdata.IRData, deviceID device.ID) *Button {
	return &Button{
		ID:       id,
		Name:     name,
		Tag:      tag,
		IRData:   irData,
		DeviceID: deviceID,
	}
}

func Factory(name Name, tag Tag) *Button {
	return New(NewID(), name, tag, nil, "")
}

func Restore(id ID, name Name, tag Tag, irData irdata.IRData) *Button {
	return &Button{
		ID:     id,
		Name:   name,
		Tag:    tag,
		IRData: irData,
	}
}

func (b *Button) LearnIR(irData irdata.IRData) {
	b.IRData = irData
}
