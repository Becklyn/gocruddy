package gocruddy

type Error struct {
	err          error
	responseCode int
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) Unwrap() error {
	return e.err
}

func NewError(responseCode int, err error) Error {
	return Error{
		err:          err,
		responseCode: responseCode,
	}
}
