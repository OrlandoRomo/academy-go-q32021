package repository

import (
	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/api"
)

type urbanDictionaryRepository struct {
	urbanDictionaryClient *api.UrbanDictionary
}

func NewUrbanDictionaryRepository(urbanDictionary *api.UrbanDictionary) *urbanDictionaryRepository {
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

func (u *urbanDictionaryRepository) GetDefinitionById(id string) (*model.List, error) {
	definitions, err := u.urbanDictionaryClient.GetDefinitionById(id)
	if err != nil {
		return nil, err
	}
	return definitions, err
}

func (u *urbanDictionaryRepository) GetConcurrentDefinitions(idType string, taskSize, workers int) (*model.List, error) {
	definitions, err := u.urbanDictionaryClient.GetConcurrentDefinitions(idType, taskSize, workers)
	if err != nil {
		return nil, err
	}
	return definitions, nil
}
