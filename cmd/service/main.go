package main

import (
	"net/http"
	"os"

	"github.com/gokultp/gstreamer/internal/dbmodels"
	"github.com/gokultp/gstreamer/internal/routes"
)

const envPort = "PORT"

func main() {
	port := os.Getenv(envPort)
	if port == "" {
		port = "8080"
	}
	err := dbmodels.InitDBConnection()
	if err != nil {
		panic(err)
	}
	router := routes.InitRoutes()
	http.ListenAndServe(":"+port, router)
}
