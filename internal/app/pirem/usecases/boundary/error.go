package boundary

type ErrCode int

const (
	CodeNotExist ErrCode = iota
	CodeAlreadyExists
	CodeInvaildInput
	CodeTimeout
	CodeBusy
	CodeDevice
	CodeInternal
	CodeNotSupported
)

type Error struct {
	Code ErrCode
	Err  error
}

func WrapErr(code ErrCode, err error) error {
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
