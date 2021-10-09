package presenter

import (
	"time"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/presenter"
)

var (
	UrbanLayout = "2006-01-02T15:04:05.999Z"
	UserLayout  = "Mon, 02-January-2006 15:04"
)

type definitionPresenter struct{}

func NewDefinitionPresenter() presenter.DefinitionPresenter {
	return &definitionPresenter{}
}

// ResponseDefinitions return the list of definitions fulfilling the DefinitionPresenter interface
func (l *definitionPresenter) ResponseDefinitions(definitionsList *model.List) *model.List {
	for _, definition := range definitionsList.Definitions {
		writtenParsed, err := time.Parse(UrbanLayout, definition.WrittenOn)
		if err != nil {
			continue
		}
		definition.WrittenOn = writtenParsed.Format(UserLayout)
	}
	return definitionsList
}
