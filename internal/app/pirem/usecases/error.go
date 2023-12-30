package usecases

type ErrCode int

const (
	CodeInvaildInput ErrCode = iota
	CodeTimeout
	CodeBusy
	CodeUnknown
	CodeDeviceNotFound
	CodeNotSupported

	CodeNotFound
	CodeAlreadyExists
	CodeDataBase
)

type Error struct {
	Code ErrCode
	Err  error
}

func WrapError(code ErrCode, err error) error {
	if err == nil {
		return nil
	}
	return &Error{
		Code: code,
		Err:  err,
	}
}

func (e *Error) Error() string {
	return e.Err.Error()
}
