package error

type IError interface {
	Error() string
	Code() int
	HTTPStatus() int
}
