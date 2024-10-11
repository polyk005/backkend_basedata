package router

import (
	"example.com/m/v2/api/controllers/book"
	"github.com/gorilla/mux"
)

type Router struct {
	BookRoutes *book.ControllerRoute
}

func NewRouter(bookRoutes *book.ControllerRoute) *Router {
	return &Router{
		BookRoutes: bookRoutes,
	}
}

func (routes *Router) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = routes.BookRoutes.Route(router)
	return router
}
