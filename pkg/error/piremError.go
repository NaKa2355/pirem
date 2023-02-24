package pirem_error

import "fmt"

var ErrDeviceNotFound = fmt.Errorf("device not found")
var ErrInvaildArgument = fmt.Errorf("invaild argument")
var ErrDeviceInternal = fmt.Errorf("device internal error")
