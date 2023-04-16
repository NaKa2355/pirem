package boundary

type ReceiveIROutput struct {
	IRData IRData
}

type GetDevicesInfoOutput struct {
	Devices []DeviceInfo
}

type GetDeviceInfoOutput struct {
	Device DeviceInfo
}

type GetIsBusyOutput struct {
	IsBusy bool
}

type GetStatusOutput struct {
	Status Status
}
