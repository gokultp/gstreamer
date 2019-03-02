package models

const (
	// TableUser is the db table for User
	TableUser = "user"
)

// User is the gorm model for user object
type User struct {
	ID          *int64  `gorm:"id"`
	Name        *string `gorm:"name"`
	Email       *string `gorm:"email"`
	DisplayName *string `gorm:"display_name"`
	Logo        *string `gorm:"logo"`
	FavStreamer *string `gor:"fav_streamer"`
}


func NewUser(id name, email, displayName, logo) (*User, error)