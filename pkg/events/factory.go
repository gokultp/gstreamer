package events

func NewEvent(etype string, streamer uint64) IEvent {
	switch etype {
	case EventFollows:
		return NewFollow(streamer)
	case EventNewFollower:
		return NewNFollower(streamer)
	}
	return NewFollow(streamer)
}
