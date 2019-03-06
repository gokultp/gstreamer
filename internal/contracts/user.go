package contracts

import (
	"strconv"

	"github.com/gokultp/gstreamer/internal/dbmodels"
)

// User is the json model for user object
type User struct {
	ID              *uint64 `json:"_id"`
	Name            *string `json:"name"`
	Email           *string `json:"email"`
	DisplayName     *string `json:"display_name"`
	Logo            *string `json:"logo"`
	FavStreamer     *uint64 `json:"fav_streamer"`
	FavStreamerName *string `json:"fav_streamer_name"`
}

func ConvertUser(user *dbmodels.User) *User {
	return &User{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		DisplayName:     user.DisplayName,
		Logo:            user.Logo,
		FavStreamer:     user.FavStreamer,
		FavStreamerName: user.FavStreamerName,
	}
}

// Streamer is the json model for Streamer user object
type Streamer struct {
	ID              *string `json:"_id"`
	Name            *string `json:"name"`
	Email           *string `json:"email"`
	DisplayName     *string `json:"display_name"`
	Logo            *string `json:"logo"`
	FavStreamer     *uint64 `json:"fav_streamer"`
	FavStreamerName *string `json:"fav_streamer_name"`
}

func ConvertStreamer(user Streamer) *User {
	id, _ := strconv.ParseInt(*user.ID, 10, 64)
	uid := uint64(id)
	return &User{
		ID:          &uid,
		Name:        user.Name,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		Logo:        user.Logo,
		FavStreamer: user.FavStreamer,
	}
}
