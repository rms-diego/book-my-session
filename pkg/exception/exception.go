package exception

type exception struct {
	message string
	status  int
}

type Exception interface {
	Error() string
	Code() int
}

func NewCustomError(message string, status int) Exception {
	return &exception{message, status}
}

func (e *exception) Error() string {
	return e.message
}

func (e *exception) Code() int {
	return e.status
}
