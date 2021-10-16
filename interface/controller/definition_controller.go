package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/api"
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
	idType := params["type"]
	items := params["items"]
	itemsPerWorker := params["items_per_workers"]
	if strings.ToLower(idType) != api.Odd && strings.ToLower(idType) != api.Even {
		model.EncodeError(w, model.ErrInvalidData{Field: "type"})
		return
	}

	itemsResponse, err := valideRange(items, "items")
	if err != nil {
		model.EncodeError(w, err)
		return
	}
	perWorker, err := valideRange(itemsPerWorker, "items_per_workers")
	if err != nil {
		model.EncodeError(w, err)
		return
	}
	definitions, err := l.definitionInteractor.GetConcurrent(idType, itemsResponse, perWorker)
	if err != nil {
		model.EncodeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(&definitions)
}

func valideRange(value, name string) (int, error) {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	if val < 0 {
		return 0, model.ErrInvalidData{Field: name}
	}
	return val, nil
}
