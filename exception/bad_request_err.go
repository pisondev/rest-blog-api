package exception

type BadRequestError struct {
	Err string
}

func NewBadRequestError(err string) BadRequestError {
	return BadRequestError{Err: err}
}
