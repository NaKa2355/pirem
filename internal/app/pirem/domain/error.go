package domain

type ErrCode int

const (
	CodeInvaildInput ErrCode = iota
	CodeInvaildOperation
	CodeNotSupported
	CodeDeviceBusy
	CodeDeviceNotFound
)

type Error struct {
	Err  error
	Code ErrCode
}

func WrapError(code ErrCode, err error) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func (e Error) Error() string {
	return e.Err.Error()
}
