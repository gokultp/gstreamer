package events

import (
	"fmt"
	"os"

	"github.com/gokultp/gstreamer/internal/helpers"
	"github.com/gokultp/gstreamer/pkg/errors"
)

const (
	EventFollows = "follows"
	envHost      = "HOST"
)

type Follow struct {
	From     uint64 `json:"from_id"`
	To       uint64 `json:"to_id"`
	FromName string `json:"from_name"`
	ToName   string `json:"to_name"`
}

func NewFollow(streamer uint64) *Follow {
	return &Follow{
		From: streamer,
	}
}

func (f *Follow) Topic() string {
	return fmt.Sprintf("https://api.twitch.tv/helix/users/follows?first=1&from_id=%d", f.From)
}

func (f *Follow) Message() string {
	return fmt.Sprintf("%s Followed %s", f.FromName, f.ToName)
}

func (f *Follow) Event() string {
	return EventFollows
}

func (f *Follow) CallbackURL() string {
	return fmt.Sprintf("%s/hooks/streamer/%d/events/%s", os.Getenv(envHost), f.From, f.Event())
}

func (f *Follow) Subscribe(token string) errors.IError {
	return helpers.SubscribeEvent(token, f.CallbackURL(), f.Topic())
}
