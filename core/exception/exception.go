package exception

type Exception struct {
	msg string
	err error
}

func (e *Exception) Error() string {
	return e.msg
}

func (e *Exception) GetError() error {
	return e.err
}

func NewException(msg string, err ...error) *Exception {
	if len(err) > 0 {
		return &Exception{
			msg: msg,
			err: err[0],
		}
	}
	return &Exception{
		msg: msg,
	}
}
