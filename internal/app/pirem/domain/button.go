package domain

type Button struct {
	ID        ButtonID
	Name      ButtonName
	Tag       ButtonTag
	HasIRData bool
}

func NewButton(id ButtonID, name ButtonName, tag ButtonTag) *Button {
	return &Button{
		ID:        id,
		Name:      name,
		Tag:       tag,
		HasIRData: false,
	}
}

func ButtonFactory(name ButtonName, tag ButtonTag, deviceID DeviceID) *Button {
	return NewButton(NewButtonID(), name, tag)
}
