package repository

import (
	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/api"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/repository"
)

type urbanDictionaryRepository struct {
	urbanDictionaryClient *api.UrbanDictionary
}

func NewUrbanDictionaryRepository(urbanDictionary *api.UrbanDictionary) repository.UrbanDictionaryRepository {
	return &urbanDictionaryRepository{urbanDictionary}
}

func (u *urbanDictionaryRepository) GetDefinitionsByTerm(term string) (*model.List, error) {
	definitionsList, err := u.urbanDictionaryClient.GetDefinitions(term)
	if err != nil {
		return nil, err
	}
	err = u.urbanDictionaryClient.Write(definitionsList)
	if err != nil {
		return nil, err
	}
	return definitionsList, err
}
func (u *urbanDictionaryRepository) GetDefinitionsFromCSV(id string) (*model.List, error) {
	definitions, err := u.urbanDictionaryClient.GetDefinitionsCSV(id)
	if err != nil {
		return nil, err
	}
	return definitions, err
}
