package error

import "fmt"

// Error encapsulates the error function
type Error struct {
	message    string
	code       int
	httpStatus int
}

// NewError returns a factory for Error
func NewError(msg string, code, status int) func(string, ...interface{}) IError {
	return func(format string, msgs ...interface{}) IError {
		return newError(msg, code, status).AddError(format, msgs)
	}
}

func newError(msg string, code, status int) *Error {
	return &Error{
		message:    msg,
		code:       code,
		httpStatus: status,
	}
}

// Error returns the error message
func (e *Error) Error() string {
	return e.message
}

// AddError adds formatted error message
func (e *Error) AddError(format string, msgs ...interface{}) *Error {
	e.message += fmt.Sprintf(format, msgs...)
	return e
}

// Code returns the error status
func (e *Error) Code() int {
	return e.code
}

// HTTPStatus returns the http status to be returned for the error
func (e *Error) HTTPStatus() int {
	return e.httpStatus
}

// String implements the stringer interace
func (e Error) String() string {
	return fmt.Sprintf("Error: %s (%d=>%d)", e.message, e.code, e.httpStatus)
}
