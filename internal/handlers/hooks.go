package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/gokultp/gstreamer/internal/ws"
	"github.com/gorilla/mux"
)

func HookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		VerifyHook(w, r)
		return
	case http.MethodPost:
		ProcessEvent(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func VerifyHook(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query()["hub.challenge"]
	if len(challenge) == 0 {
		w.WriteHeader(400)
		return
	}
	w.Write([]byte(challenge[0]))
}

func ProcessEvent(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	defer r.Body.Close()
	event, _ := ioutil.ReadAll(r.Body)
	ws.WriteEvent(vals["id"], event)
}
