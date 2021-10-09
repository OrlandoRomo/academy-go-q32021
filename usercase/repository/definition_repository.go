package repository

import (
	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
)

type UrbanDictionaryRepository interface {
	GetDefinitionsByTerm(term string) (*model.List, error)
	GetDefinitionById(id string) (*model.List, error)
	GetConcurrentDefinitions() (*model.List, error)
}
