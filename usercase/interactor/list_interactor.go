package interactor

import (
	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/presenter"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/repository"
)

type listInteractor struct {
	ListRepository repository.ListRepository
	ListPresenter  presenter.ListPresenter
}

type ListInteractor interface {
	Get(term string) ([]*model.List, error)
	GetFromCSV() ([]*model.List, error)
}

func NewListInteractor(repository repository.ListRepository, presenter presenter.ListPresenter) *listInteractor {
	return &listInteractor{repository, presenter}
}

func (l *listInteractor) Get(term string) ([]*model.List, error) {
	definitions, err := l.ListRepository.FindDefinitionsList(term)
	if err != nil {
		return nil, err
	}
	return l.ListPresenter.ResponseList(definitions), nil
}

func (l *listInteractor) GetFromCSV() ([]*model.List, error) {
	return nil, nil
}
