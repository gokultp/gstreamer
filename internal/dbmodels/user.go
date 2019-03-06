package dbmodels

import (
	"github.com/gokultp/gstreamer/internal/serviceerrors"
	"github.com/gokultp/gstreamer/pkg/errors"
)

const (
	// TableUser is the db table for User
	TableUser = "user"
)

// User is the gorm model for user object
type User struct {
	ID              *uint64 `gorm:"column:id;primary_key"`
	Name            *string `gorm:"column:name"`
	Email           *string `gorm:"column:email"`
	DisplayName     *string `gorm:"column:display_name"`
	Logo            *string `gorm:"column:logo"`
	FavStreamer     *uint64 `gorm:"column:fav_streamer"`
	FavStreamerName *string `gorm:"column:fav_streamer_name"`
	AccessToken     *string `gorm:"column:access_token"`
	RefreshToken    *string `gorm:"column:refresh_token"`
}

func NewUser(id uint64, name, email, displayName, logo string) *User {
	return &User{
		ID:          &id,
		Name:        &name,
		Email:       &email,
		DisplayName: &displayName,
		Logo:        &logo,
	}
}

func GetUserByID(id uint64) (*User, errors.IError) {
	user := User{}
	if err := Connection.Where("id=?", id).First(&user).Error; err != nil {
		return nil, serviceerrors.DBFetchError(err.Error())
	}
	return &user, nil
}

func (u *User) CreateUser() errors.IError {
	if err := Connection.Create(u).Error; err != nil {
		return serviceerrors.DBUpdateError(err.Error())
	}
	return nil
}

func (u *User) UpdateUser() errors.IError {
	if err := Connection.Model(u).Updates(u).Error; err != nil {
		return serviceerrors.DBUpdateError(err.Error())
	}
	return nil
}
