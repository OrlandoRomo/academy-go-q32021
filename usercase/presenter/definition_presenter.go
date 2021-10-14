package presenter

import "github.com/OrlandoRomo/academy-go-q32021/domain/model"

type DefinitionPresenter interface {
	ResponseDefinitions(definitionsList *model.List) (*model.List, error)
}
