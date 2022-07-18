package dto

type Error struct {
	Code    int         `json:"error_code"`
	Message string      `json:"error_msg"`
	Err     interface{} `json:"error"`
}

type Option func(*Error)

func NewError(err error, opts ...Option) Error {
	e := Error{
		Code:    0,
		Message: err.Error(),
		Err:     err,
	}

	for _, opt := range opts {
		opt(&e)
	}

	return e
}

func SetCode(c int) Option {
	return func(err *Error) {
		err.Code = c
	}
}

func SetMessage(msg string) Option {
	return func(err *Error) {
		err.Message = msg
	}
}
