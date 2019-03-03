package serviceerrors

import (
	errlib "github.com/gokultp/gstreamer/pkg/errors"
)

var (
	// DBConectionError is the error thrown if there is an error while connecting db.
	DBConectionError = errlib.NewError("Error while connecting to db: ", 1001, 500)
	DBFetchError     = errlib.NewError("Error while fetching data from db: ", 1002, 500)
	DBUpdateError    = errlib.NewError("Error while updating data in db: ", 1003, 500)
)
