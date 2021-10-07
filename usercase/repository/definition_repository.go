package repository

import (
	"os"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
)

type UrbanDictionaryRepository interface {
	GetDefinitionsByTerm(term string) (*model.List, error)
	GetDefinitionsFromCSV(id string) (*model.List, error)
}

type UrbanReaderWriter interface {
	Read() (*os.File, error)
	Write(definitionsList *model.List) error
}
