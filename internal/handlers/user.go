package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gokultp/gstreamer/internal/contracts"
	"github.com/gokultp/gstreamer/internal/helpers"
	"github.com/gokultp/gstreamer/internal/serviceerrors"
	"github.com/gokultp/gstreamer/pkg/errors"
	"github.com/gokultp/gstreamer/pkg/events"

	"github.com/gokultp/gstreamer/internal/dbmodels"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUserByID(w, r)
		return
	case http.MethodPost:
		SetFavStreamer(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		HandleError(err, w)
		return
	}

	user, _ := dbmodels.GetUserByID(*userID)
	if user == nil {
		HandleError(serviceerrors.ResourceNotFoundError(""), w)
		return
	}
	jsonResponse(contracts.ConvertUser(user), w)
}

func SetFavStreamer(w http.ResponseWriter, r *http.Request) {
	userID, ierr := CheckAuthenticated(r)
	if ierr != nil {
		HandleError(ierr, w)
		return
	}
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(serviceerrors.BadRequestError("%v", err), w)
		return
	}
	var user contracts.User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		HandleError(serviceerrors.BadRequestError("%v", err), w)
		return
	}
	dbuser, ierr := dbmodels.GetUserByID(*userID)
	if ierr != nil {
		HandleError(ierr, w)
		return
	}
	streamer, ierr := helpers.GetTwitchUserByName(*user.FavStreamerName, *dbuser.AccessToken)
	if ierr != nil {
		HandleError(ierr, w)
		return
	}
	dbuser.FavStreamer = streamer.ID
	dbuser.FavStreamerName = user.FavStreamerName
	ierr = dbuser.UpdateUser()
	if ierr != nil {
		HandleError(ierr, w)
		return
	}
	ierr = subscribeEvents(*streamer.ID, *dbuser.AccessToken, events.EventFollows, events.EventNewFollower)
	if ierr != nil {
		HandleError(ierr, w)
		return
	}
	jsonResponse(contracts.ConvertUser(dbuser), w)
}

func getUserID(r *http.Request) (*uint64, errors.IError) {
	userID, err := CheckAuthenticated(r)
	if err != nil {

		return nil, err
	}
	return userID, nil
}

func subscribeEvents(streamerId uint64, accessToken string, eventIds ...string) errors.IError {
	for _, ev := range eventIds {
		event := events.NewEvent(ev, streamerId)
		if err := event.Subscribe(accessToken); err != nil {
			return err
		}
	}
	return nil
}
