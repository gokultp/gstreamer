package dbmodels

import (
	"fmt"
	"os"

	"github.com/gokultp/gstreamer/internal/serviceerrors"
	"github.com/gokultp/gstreamer/pkg/errors"

	"gopkg.in/jinzhu/gorm.v1"
)

const (
	envDBHost     = "DB_HOST"
	envDBPort     = "DB_PORT"
	envDBUser     = "DB_USER"
	envDBPassword = "DB_PASSWORD"
	envDBNAME     = "DB_NAME"
)

// Connection is the singleton maintained for db connection
var Connection gorm.DB

// InitDBConnection will initialises a db connection
func InitDBConnection() errors.IError {
	host := os.Getenv(envDBHost)
	port := os.Getenv(envDBPort)
	user := os.Getenv(envDBUser)
	password := os.Getenv(envDBPassword)
	dbName := os.Getenv(envDBNAME)
	Connection, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbName, password))
	if err != nil {
		return serviceerrors.DBConectionError(err.Error())
	}
	return nil

}
