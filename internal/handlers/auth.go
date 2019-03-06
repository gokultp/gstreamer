package handlers

import (
	"net/http"

	"github.com/gokultp/gstreamer/internal/dbmodels"

	"github.com/gokultp/gstreamer/internal/helpers"
	"github.com/gokultp/gstreamer/internal/serviceerrors"
)

const (
	queryFieldCode = "code"
)

func AuthCBHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		CheckAuth(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query()[queryFieldCode]
	if len(code) == 0 || code[0] == "" {
		HandleError(serviceerrors.BadRequestError("No code found in query string"), w)
		return
	}
	token, err := helpers.GetAuthToken(code[0])
	if err != nil {
		HandleError(err, w)
		return
	}
	user, err := helpers.GetUserInfo(*token.AccessToken)
	if err != nil {
		HandleError(err, w)
		return
	}
	userModel, err := dbmodels.GetUserByID(*user.ID)
	if userModel != nil {
		userModel.AccessToken = token.AccessToken
		userModel.RefreshToken = token.RefreshToken
		err = userModel.UpdateUser()
		if err != nil {
			HandleError(err, w)
			return
		}
		SetSession(*user.ID, &w)
		http.Redirect(w, r, "/", 301)
		return
	}
	userModel = dbmodels.NewUser(*user.ID, *user.Name, *user.Email, *user.DisplayName, *user.Logo)
	userModel.AccessToken = token.AccessToken
	userModel.RefreshToken = token.RefreshToken
	err = userModel.CreateUser()
	if err != nil {
		HandleError(err, w)
		return
	}
	SetSession(*user.ID, &w)
	http.Redirect(w, r, "/", 301)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		AuthRedirect(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func AuthRedirect(w http.ResponseWriter, r *http.Request) {
	authURL := helpers.GetAuthRedirectURL()
	http.Redirect(w, r, authURL, 301)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ResetSession(&w)
	http.Redirect(w, r, "/", 301)
}
