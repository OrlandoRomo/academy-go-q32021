package repository

import "github.com/OrlandoRomo/academy-go-q32021/domain/model"

type ListRepository interface {
	FindDefinitionsList(term string) ([]*model.List, error)
}
