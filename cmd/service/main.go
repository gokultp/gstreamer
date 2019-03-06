package main

import (
	"net/http"

	"github.com/gokultp/gstreamer/internal/dbmodels"
	"github.com/gokultp/gstreamer/internal/routes"
)

func main() {

	err := dbmodels.InitDBConnection()
	if err != nil {
		panic(err)
	}
	router := routes.InitRoutes()
	http.ListenAndServe(":8080", router)
}
