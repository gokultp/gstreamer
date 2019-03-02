package errors

type IError interface {
	Error() string
	Code() int
	HTTPStatus() int
}
