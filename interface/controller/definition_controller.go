package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/interactor"
	"github.com/gorilla/mux"
)

var ErrTermEmpty = errors.New("term should not be empty")
var ErrReadingCSV = errors.New("there was a problem reading the CSV file")
var ErrCreatingCSV = errors.New("there was a problem creating the CSV file")
var ErrWritingCSV = errors.New("there was a problem writing the CSV file")

type definitionController struct {
	definitionInteractor interactor.DefinitionInteractor
}

type DefinitionController interface {
	GetDefinitions(w http.ResponseWriter, r *http.Request)
	GetDefinitionsFromCSV(w http.ResponseWriter, r *http.Request)
}

func NewDefinitionController(r interactor.DefinitionInteractor) DefinitionController {
	return &definitionController{r}
}

// GetDefinitions handles the requests and responses of the /definitions/ endpoint
func (l *definitionController) GetDefinitions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	term, ok := params["term"]
	if !ok || term == "" {
		model.NewHttpError(w, http.StatusBadRequest, ErrTermEmpty.Error())
		return
	}
	definitions, err := l.definitionInteractor.Get(term)
	if err != nil {
		model.NewHttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(&definitions)
}

// GetDefinitionsFromCSV handles the requests and responses of the /definitions/csv/ endpoint
func (l *definitionController) GetDefinitionsFromCSV(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	concurrentStr := params["concurrent"]
	concurrent, err := strconv.ParseBool(concurrentStr)
	if err != nil {
		concurrent = false
	}
	if concurrent {
		// do your concurrent thing
	}
	definitions, err := l.definitionInteractor.GetFromCSV(params["id"])
	if err != nil {
		model.NewHttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(&definitions)
}
