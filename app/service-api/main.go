package serviceapi

import (
	"can-i-go/app/service-api/config"
	"can-i-go/app/service-api/webserver"
)

type API struct {
	srv    webserver.Server
	Config config.Config
}

func New() API {
	return API{
		srv: webserver.Server{},
		Config: config.Config{
			Addr: "0.0.0.0",
			Port: 4000,
		},
	}
}

func (api *API) Start() {
	api.srv = webserver.New(api.Config)
	api.srv.Start(BuildPipeline)
}
