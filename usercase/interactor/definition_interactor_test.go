package interactor

import (
	"testing"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockDefinitionsRepo struct {
	mock.Mock
}

func (m MockDefinitionsRepo) Get(term string) (*model.List, error) {
	args := m.Called()
	return args.Get(0).(*model.List), args.Error(1)
}

func (m MockDefinitionsRepo) GetFromCSV(id string) (*model.List, error) {
	args := m.Called()
	return args.Get(0).(*model.List), args.Error(1)
}

func TestGet(t *testing.T) {
	t.Run("Something", func(t *testing.T) {
		mock := MockDefinitionsRepo{}
		mock.On("Get").Return()
	})
}

func TestGetFromCSV(t *testing.T) {

}
