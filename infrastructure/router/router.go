package router

import (
	"net/http"

	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/middleware"
	"github.com/OrlandoRomo/academy-go-q32021/interface/controller"
	"github.com/gorilla/mux"
)

func NewRouter(c controller.AppController) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.HeadersMiddleware)

	router = router.PathPrefix("/api/v1/").Subrouter()

	router.HandleFunc("/definitions/", c.Definitions.GetDefinitions).
		Queries("term", "{term}").
		Methods(http.MethodGet)

	router.HandleFunc("/definitions/csv/{id:[0-9a-zA-Z\\W]+|}/", c.Definitions.GetDefinitionsFromCSV).
		Queries("concurrent", "{concurrent:(?:true|false)}").
		Methods(http.MethodGet)
	return router
}
