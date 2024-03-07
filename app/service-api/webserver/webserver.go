package webserver

import (
	"can-i-go/app/service-api/config"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Server http.Server
}

type PipelineBuilder func(*Server, *mux.Router)

func New(cfg config.Config) Server {
	return Server{
		Server: http.Server{
			Addr: fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
		},
	}
}

func (srv *Server) Start(builder PipelineBuilder) {
	r := mux.NewRouter()
	builder(srv, r)
	http.Handle("/", r)
	log.Printf("Listening on %s\n", srv.Server.Addr)
	if err := srv.Server.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
