package repository

import "github.com/OrlandoRomo/academy-go-q32021/domain/model"

type listRepository struct {
	model.UrbanDictionary
}

func NewListRepository(apiKey string) *listRepository {
	return &listRepository{
		UrbanDictionary: *model.NewUrbanDictionary(apiKey),
	}
}

func (l *listRepository) FindDefinitionsList(term string) ([]*model.List, error) {
	definitions, err := l.UrbanDictionary.GetDefinitionsByTerm(term)
	if err != nil {
		return nil, err
	}
	return definitions, nil
}
