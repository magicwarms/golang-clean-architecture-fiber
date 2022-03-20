package shared

// errorString is a trivial implementation of error.
type SystemError struct {
	s string
}

func NewSystemError(text string) error {
	return &SystemError{text}
}

func (e *SystemError) Error() string {
	return "maaf gan system lagi error"
}
