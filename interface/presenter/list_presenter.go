package presenter

import (
	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/presenter"
)

type listPresenter struct{}

func NewListPresenter() presenter.ListPresenter {
	return &listPresenter{}
}

func (l *listPresenter) ResponseList(definitions []*model.List) []*model.List {
	return definitions
}
