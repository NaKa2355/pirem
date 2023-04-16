package boundary

type ErrCode int

const (
	CodeNotExist ErrCode = iota
	CodeInvaildInput
	CodeTimeout
	CodeBusy
	CodeDevice
)

type Error struct {
	Code ErrCode
	Err  error
}

func Wrap(code ErrCode, err error) error {
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
