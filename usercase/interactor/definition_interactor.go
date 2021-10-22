package interactor

import (
	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/presenter"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/repository"
)

type DefinitionInteractor interface {
	Get(term string) (*model.List, error)
	GetFromCSV(id string) (*model.List, error)
	GetConcurrent(idType string, taskSize, perWorker int) (*model.List, error)
}

type definitionInteractor struct {
	urbanDictionaryRepository repository.UrbanDictionaryRepository
	definitionsPresenter      presenter.DefinitionPresenter
}

func NewDefinitionInteractor(repository repository.UrbanDictionaryRepository, presenter presenter.DefinitionPresenter) *definitionInteractor {
	return &definitionInteractor{repository, presenter}
}

func (d *definitionInteractor) Get(term string) (*model.List, error) {
	definitionsList, err := d.urbanDictionaryRepository.GetDefinitionsByTerm(term)
	if err != nil {
		return nil, err
	}
	return d.definitionsPresenter.ResponseDefinitions(definitionsList)
}

func (d *definitionInteractor) GetFromCSV(id string) (*model.List, error) {
	definitionsList, err := d.urbanDictionaryRepository.GetDefinitionById(id)
	if err != nil {
		return nil, err
	}
	return d.definitionsPresenter.ResponseDefinitions(definitionsList)
}

func (d *definitionInteractor) GetConcurrent(idType string, taskSize, perWorker int) (*model.List, error) {
	definitionsList, err := d.urbanDictionaryRepository.GetConcurrentDefinitions(idType, taskSize, perWorker)
	if err != nil {
		return nil, err
	}
	return d.definitionsPresenter.ResponseDefinitions(definitionsList)
}
