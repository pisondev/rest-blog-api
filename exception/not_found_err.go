package exception

type NotFoundError struct {
	Err string
}

func NewNotFoundError(err string) NotFoundError {
	return NotFoundError{Err: err}
}
