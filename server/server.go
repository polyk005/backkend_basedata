package server

import (
	"net/http"
	"src/api/router"
	"src/internal/config"

	"github.com/codegangsta/negroni"
	"github.com/docker/docker/api/server/router"
	"github.com/docker/docker/daemon/config"
	"github.com/fasthttp/router"
)

type Server struct {
	AppConfig *config.Config
	Router    *router.Router
}

func NewServer(appConfig *config.Config, router *router.Router) *Server {
	return &Server{
		AppConfig: appConfig,
		Router:    router,
	}
}

// Run server
func (server *Server) Run() {
	ngRouter := server.Router.InitRoutes()
	ngClassic := negroni.Classic()
	ngClassic.UseHandler(ngRouter)
	err := http.ListenAndServe(":8080", ngClassic)
	if err != nil {
		return
	}
}
