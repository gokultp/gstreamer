package dbmodels

import (
	"github.com/gokultp/gstreamer/pkg/errors"
)

const (
	// TableUser is the db table for User
	TableUser = "user"
)

// User is the gorm model for user object
type User struct {
	ID          *int64  `gorm:"column:id"`
	Name        *string `gorm:"column:name"`
	Email       *string `gorm:"column:email"`
	DisplayName *string `gorm:"column:display_name"`
	Logo        *string `gorm:"column:logo"`
	FavStreamer *string `gorm:"column:fav_streamer"`
}

func NewUser(id, name, email, displayName, logo string) (*User, errors.IError) {

}

func GetUserByID(id string) (*User, errors.IError) {

}
