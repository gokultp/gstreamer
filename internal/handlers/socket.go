package handlers

import (
	"net/http"

	"github.com/gokultp/gstreamer/internal/serviceerrors"

	"github.com/gokultp/gstreamer/internal/dbmodels"
	"github.com/gokultp/gstreamer/internal/ws"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SockHandler(w http.ResponseWriter, r *http.Request) {
	userID, ierr := CheckAuthenticated(r)
	if ierr != nil {
		HandleError(ierr, w)
		return
	}
	user, ierr := dbmodels.GetUserByID(*userID)
	if ierr != nil {
		HandleError(ierr, w)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		HandleError(serviceerrors.DBConectionError(err.Error()), w)
		return
	}
	ws.StartEventListener(*user.FavStreamer, conn)

}
