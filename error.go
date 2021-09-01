package gocruddy

type CRUDError struct {
	err          error
	responseCode int
}

func (e CRUDError) Error() string {
	return e.err.Error()
}

func (e CRUDError) Unwrap() error {
	return e.err
}

func NewCRUDError(responseCode int, err error) CRUDError {
	return CRUDError{
		err:          err,
		responseCode: responseCode,
	}
}
