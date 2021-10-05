package presenter

import "github.com/OrlandoRomo/academy-go-q32021/domain/model"

type ListPresenter interface {
	ResponseList([]*model.List) []*model.List
}