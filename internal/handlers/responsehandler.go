package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gokultp/gstreamer/internal/serviceerrors"
	"github.com/gokultp/gstreamer/pkg/errors"

	"github.com/gorilla/securecookie"
)

const (
	key64      = "_afwDY&nNHAoei^@HTz+k)e=mm8SEJpT(3V5wyo+8+W*(VD-XWMCQdF100q@IBQ!"
	key32      = "pXCEF5NL#SKu6_9xNOh^sJsp@WQb_gp#"
	cookieName = "gs_session"
)

func jsonResponse(data interface{}, w http.ResponseWriter) {
	btData, _ := json.Marshal(data)
	w.Write(btData)
}

var cookieHandler = securecookie.New([]byte(key64), []byte(key32))

func SetSession(userId uint64, response *http.ResponseWriter) {
	value := map[string]uint64{
		"id": userId,
	}
	if encoded, err := cookieHandler.Encode(cookieName, value); err == nil {
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(*response, cookie)
	}
}
func ResetSession(response *http.ResponseWriter) {

	exp := time.Now().Add(-60 * 24 * time.Hour)
	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   "",
		Path:    "/",
		Expires: exp,
	}
	http.SetCookie(*response, cookie)

}
func CheckAuthenticated(request *http.Request) (*uint64, errors.IError) {

	cookie, err := request.Cookie(cookieName)
	if err != nil {
		return nil, serviceerrors.UnAuthorized(err.Error())
	}

	cookieValue := make(map[string]uint64)

	err = cookieHandler.Decode(cookieName, cookie.Value, &cookieValue)
	if err != nil {

		return nil, serviceerrors.UnAuthorized(err.Error())
	}
	id := cookieValue["id"]
	return &id, nil

}
