package interactor

import (
	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/repository"
)

var CSVPath = "definitions.csv"

type DefinitionInteractor interface {
	Get(term string) (*model.List, error)
	GetFromCSV(id string) (*model.List, error)
}

type definitionInteractor struct {
	urbanDictionaryRepository repository.UrbanDictionaryRepository
}

func NewDefinitionInteractor(repository repository.UrbanDictionaryRepository) *definitionInteractor {
	return &definitionInteractor{repository}
}

func (d *definitionInteractor) Get(term string) (*model.List, error) {
	definitionsList, err := d.urbanDictionaryRepository.GetDefinitionsByTerm(term)
	if err != nil {
		return nil, err
	}
	return definitionsList, nil
}
func (d *definitionInteractor) GetFromCSV(id string) (*model.List, error) {
	definitionsList, err := d.urbanDictionaryRepository.GetDefinitionsFromCSV(id)
	if err != nil {
		return nil, err
	}
	return definitionsList, nil
}
