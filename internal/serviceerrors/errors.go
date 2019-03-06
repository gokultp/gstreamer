package serviceerrors

import (
	"net/http"

	errlib "github.com/gokultp/gstreamer/pkg/errors"
)

var (
	// DBConectionError is the error thrown if there is an error while connecting db.
	DBConectionError      = errlib.NewError("Error while connecting to db: ", 1001, http.StatusInternalServerError)
	DBFetchError          = errlib.NewError("Error while fetching data from db: ", 1002, http.StatusInternalServerError)
	DBUpdateError         = errlib.NewError("Error while updating data in db: ", 1003, http.StatusInternalServerError)
	TwitchRequestError    = errlib.NewError("Something went wrong while sending request to twitch: ", 1004, http.StatusInternalServerError)
	BadRequestError       = errlib.NewError("Bad request: ", 1005, http.StatusBadRequest)
	ResourceNotFoundError = errlib.NewError("Resource not found: ", 1006, http.StatusNotFound)
	UnAuthorized          = errlib.NewError("UnAuthorized", 1007, http.StatusUnauthorized)
)
