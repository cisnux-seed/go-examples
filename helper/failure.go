package helper

type ValidationFailure struct {
	Message string
}

func (f *ValidationFailure) Error() string {
	return f.Message
}

type NotFoundFailure struct {
	Message string
}

func (f *NotFoundFailure) Error() string {
	return f.Message
}
