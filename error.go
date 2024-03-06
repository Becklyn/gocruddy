package gocruddy

type Error struct {
	err          error
	respond      bool
	responseCode int
}

func (e Error) Respond() Error {
	e.respond = true
	return e
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
		respond:      false,
		responseCode: responseCode,
	}
}
