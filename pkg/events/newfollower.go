package events

import (
	"fmt"
	"os"

	"github.com/gokultp/gstreamer/internal/helpers"
	"github.com/gokultp/gstreamer/pkg/errors"
)

const (
	EventNewFollower = "newfollower"
)

type NFollower struct {
	From     uint64 `json:"from_id"`
	To       uint64 `json:"to_id"`
	FromName string `json:"from_name"`
	ToName   string `json:"to_name"`
}

func NewNFollower(streamer uint64) *NFollower {
	return &NFollower{
		To: streamer,
	}
}

func (f *NFollower) Topic() string {
	return fmt.Sprintf("https://api.twitch.tv/helix/users/follows?first=1&to_id=%d", f.To)
}

func (f *NFollower) Message() string {
	return fmt.Sprintf("%s Followed %s", f.ToName, f.FromName)
}

func (f *NFollower) Event() string {
	return EventNewFollower
}

func (f *NFollower) CallbackURL() string {
	return fmt.Sprintf("%s/hooks/streamer/%d/events/%s", os.Getenv(envHost), f.To, f.Event())
}

func (f *NFollower) Subscribe(token string) errors.IError {
	return helpers.SubscribeEvent(token, f.CallbackURL(), f.Topic())
}
