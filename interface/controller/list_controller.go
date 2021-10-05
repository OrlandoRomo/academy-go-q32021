package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/academy-go-q32021/usercase/interactor"
	"github.com/gorilla/mux"
)

var ErrTermEmpty = errors.New("term should not be empty")

type listController struct {
	listInteractor interactor.ListInteractor
}

type ListController interface {
	GetDefinitions(w http.ResponseWriter, r *http.Request)
	GetDefinitionsFromCSV(w http.ResponseWriter, r *http.Request)
}

func NewListController(r interactor.ListInteractor) ListController {
	return &listController{r}
}

func (l *listController) GetDefinitions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	term, ok := params["term"]
	if !ok || term == "" {
		http.Error(w, ErrTermEmpty.Error(), http.StatusBadRequest)
		return
	}
	definitions, err := l.listInteractor.Get(term)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&definitions)
}

func (l *listController) GetDefinitionsFromCSV(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	concurrentStr, ok := params["term"]
	if !ok {
		fmt.Println("term is empty or is not included")
		return
	}

	concurrent, err := strconv.ParseBool(concurrentStr)
	if err != nil {
		fmt.Println("term is empty or is not included")
		return
	}
	if concurrent {
		// do your concurrent thing
	}
}
