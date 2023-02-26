package pirem_error

import (
	"errors"
)

var ErrDeviceNotFound = errors.New("device not found")
var ErrInvaildArgument = errors.New("invaild argument")
var ErrDeviceInternal = errors.New("device internal error")
