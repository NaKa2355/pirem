package domain

type Remote struct {
	ID       RemoteID
	Name     RemoteName
	DeviceID DeviceID
	Tag      RemoteTag
	Buttons  []*Button
}

func NewRemote(id RemoteID, name RemoteName, deviceID DeviceID, tag RemoteTag, buttons []*Button) *Remote {
	return &Remote{
		ID:       id,
		Name:     name,
		Tag:      tag,
		DeviceID: deviceID,
		Buttons:  buttons,
	}
}

func RemoteFactory(name RemoteName, deviceID DeviceID, tag RemoteTag, buttons []*Button) *Remote {
	return NewRemote(NewRemoteID(), name, deviceID, tag, buttons)
}

func RestoreRemote(id RemoteID, name RemoteName, deviceID DeviceID, tag RemoteTag, buttons []*Button) *Remote {
	return &Remote{
		ID:       id,
		Name:     name,
		Tag:      tag,
		DeviceID: deviceID,
		Buttons:  buttons,
	}
}

func (r *Remote) UpdateRemote(name RemoteName, deviceID DeviceID) {
	r.Name = name
	r.DeviceID = deviceID
}
