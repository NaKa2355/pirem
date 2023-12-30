package domain

type Button struct {
	ID       ButtonID
	Name     ButtonName
	Tag      ButtonTag
	IRData   IRData
	DeviceID DeviceID
}

func New(id ButtonID, name ButtonName, tag ButtonTag, irData IRData, deviceID DeviceID) *Button {
	return &Button{
		ID:       id,
		Name:     name,
		Tag:      tag,
		IRData:   irData,
		DeviceID: deviceID,
	}
}

func Factory(name ButtonName, tag ButtonTag) *Button {
	return New(NewButtonID(), name, tag, nil, "")
}

func Restore(id ButtonID, name ButtonName, tag ButtonTag, irData IRData) *Button {
	return &Button{
		ID:     id,
		Name:   name,
		Tag:    tag,
		IRData: irData,
	}
}

func (b *Button) LearnIR(irData IRData) {
	b.IRData = irData
}
