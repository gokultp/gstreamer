package dbmodels

import (
	"github.com/gokultp/gstreamer/internal/serviceerrors"
	"github.com/gokultp/gstreamer/pkg/errors"
	"gopkg.in/jinzhu/gorm.v1"
)

const (
	// TableUser is the db table for User
	TableUser = "user"
)

// User is the gorm model for user object
type User struct {
	gorm.Model
	ID          *uint64 `gorm:"column:id"`
	Name        *string `gorm:"column:name"`
	Email       *string `gorm:"column:email"`
	DisplayName *string `gorm:"column:display_name"`
	Logo        *string `gorm:"column:logo"`
	FavStreamer *string `gorm:"column:fav_streamer"`
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
