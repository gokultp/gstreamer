package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gokultp/gstreamer/internal/handlers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/auth", handlers.AuthHandler)
	router.HandleFunc("/logout", handlers.LogoutHandler)
	router.HandleFunc("/auth/cb", handlers.AuthCBHandler)
	router.HandleFunc("/user", handlers.UserHandler)
	router.HandleFunc("/user/stream", handlers.StreamHandler)
	router.HandleFunc("/hooks/streamer/{id:[0-9]+}/events/{event:[a-z]+}", handlers.HookHandler)
	router.HandleFunc("/socket", handlers.SockHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../../web/build/")))

	return router

}
