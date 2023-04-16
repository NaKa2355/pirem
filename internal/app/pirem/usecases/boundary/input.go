package boundary

import "encoding/json"

type SendIRInput struct {
	ID     string
	IRData IRData
}

type ReceiveIRInput struct {
	ID string
}

type GetDeviceInput struct {
	ID string
}

type GetIsBusyInput struct {
	ID string
}

type GetStatusInput struct {
	ID string
}

type AddDeviceInput struct {
	Name       string
	ID         string
	PluginPath string
	Conf       json.RawMessage
}
