package router

import (
	"net/http"

	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/middleware"
	"github.com/OrlandoRomo/academy-go-q32021/interface/controller"
	"github.com/gorilla/mux"
)

// NewRouter returns a mux router with the needed resources
func NewRouter(c controller.AppController) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.HeadersMiddleware)

	router = router.PathPrefix("/api/v1/").Subrouter()

	// /api/v1/definitions/
	router.HandleFunc("/definitions/", c.Definitions.GetDefinitions).
		Queries("term", "{term}").
		Methods(http.MethodGet)

	// /api/v1/definitions/{id}/
	router.HandleFunc("/definitions/{id:[0-9a-zA-Z\\W]+|}/", c.Definitions.GetDefinitionsFromCSV).
		Methods(http.MethodGet)

	// /api/v1/definitions-csv/
	router.HandleFunc("/definitions-csv/", c.Definitions.GetConcurrentDefinitions).Queries(
		"type", "{type}",
		"items", "{items}",
		"items_per_workers", "{items_per_workers}",
	).Methods(http.MethodGet)
	return router
}
