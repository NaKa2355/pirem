package entity

type ErrCode int

const (
	CodeInvaildInput ErrCode = iota
	CodeTimeout
	CodeBusy
	CodeNotSupported
)

type Error struct {
	Code ErrCode
	Err  error
}

func WrapErr(code ErrCode, err error) error {
	if err == nil {
		return err
	}
	return &Error{
		Code: code,
		Err:  err,
	}
}

func (e *Error) Error() string {
	return e.Err.Error()
}
