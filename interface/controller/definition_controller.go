package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/interactor"
	"github.com/gorilla/mux"
)

type definitionController struct {
	definitionInteractor interactor.DefinitionInteractor
}

type DefinitionController interface {
	GetDefinitions(w http.ResponseWriter, r *http.Request)
	GetDefinitionsFromCSV(w http.ResponseWriter, r *http.Request)
	GetConcurrentDefinitions(w http.ResponseWriter, r *http.Request)
}

func NewDefinitionController(r interactor.DefinitionInteractor) DefinitionController {
	return &definitionController{r}
}

// GetDefinitions handles the requests and responses of the /definitions/ endpoint
func (l *definitionController) GetDefinitions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	term, ok := params["term"]
	if !ok || term == "" {
		model.EncodeError(w, model.ErrInvalidData{Field: "term"})
		return
	}
	definitions, err := l.definitionInteractor.Get(term)
	if err != nil {
		model.EncodeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(&definitions)
}

// GetDefinitionsFromCSV handles the requests and responses of the /definitions/csv/{id} endpoint
func (l *definitionController) GetDefinitionsFromCSV(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	definitions, err := l.definitionInteractor.GetFromCSV(params["id"])
	if err != nil {
		model.EncodeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(&definitions)
}

// GetConcurrentDefinitions handles the requests and responses of the /definitions/csv/ endpoint
func (l *definitionController) GetConcurrentDefinitions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	concurrentStr := params["concurrent"]
	concurrent, err := strconv.ParseBool(concurrentStr)
	if err != nil {
		concurrent = false
	}
	if concurrent {
		// do your concurrent thing
	}
}
