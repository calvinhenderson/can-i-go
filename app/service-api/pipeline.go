package serviceapi

import (
	"can-i-go/app/service-api/routes"
	"can-i-go/app/service-api/webserver"
	"net/http"

	"github.com/gorilla/mux"
)

func BuildPipeline(srv *webserver.Server, r *mux.Router) {
	r.HandleFunc("/ping", routes.Ping()).Methods(http.MethodGet)
}
