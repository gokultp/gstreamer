package events

import "github.com/gokultp/gstreamer/pkg/errors"

type IEvent interface {
	Topic() string
	Message() string
	Event() string
	CallbackURL() string
	Subscribe(token string) errors.IError
}
