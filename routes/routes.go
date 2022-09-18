package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	ProfileRoutes(r)
	FilmRoutes(r)
	CategoryRoutes(r)
	EpisodeRoutes(r)
	TransactionRoutes(r)
}
