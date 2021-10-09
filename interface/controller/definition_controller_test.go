package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockInteractor struct {
	mock.Mock
}

func (m MockInteractor) Get(term string) (*model.List, error) {
	args := m.Called(term)
	return args.Get(0).(*model.List), args.Error(1)
}
func (m MockInteractor) GetFromCSV(id string) (*model.List, error) {
	args := m.Called(id)
	return args.Get(0).(*model.List), args.Error(1)
}
func (m MockInteractor) GetConcurrent() (*model.List, error) {
	args := m.Called()
	return args.Get(0).(*model.List), args.Error(1)
}

func TestGetDefinitions(t *testing.T) {
	mockInteractor := MockInteractor{}
	t.Run("success - valid request", func(t *testing.T) {
		mockInteractor.On("Get", "sample").Return(&model.List{
			Definitions: []*model.Definition{
				{Word: "sample"},
			},
		}, nil)
		controller := NewDefinitionController(mockInteractor)
		req, err := http.NewRequest(http.MethodGet, "/definitions/", nil)
		assert.Nil(t, err, "new request error should be nil")
		req = mux.SetURLVars(req, map[string]string{"term": "sample"})
		rec := httptest.NewRecorder()
		controller.GetDefinitions(rec, req)
		assert.NotNil(t, rec.Body, "the body response should be not nil")
		assert.Equal(t, rec.Result().StatusCode, http.StatusOK)
	})

	t.Run("failure - bad request", func(t *testing.T) {
		mockInteractor.On("Get", "").Return(&model.List{}, model.ErrInvalidData{Field: "term"})
		controller := NewDefinitionController(mockInteractor)
		req, err := http.NewRequest(http.MethodGet, "/definitions/", nil)
		assert.Nil(t, err, "new request error should be nil")
		req = mux.SetURLVars(req, map[string]string{"term": ""})
		rec := httptest.NewRecorder()
		controller.GetDefinitions(rec, req)
		assert.NotNil(t, rec.Body, "the body response should be not nil")
		assert.Equal(t, rec.Result().StatusCode, http.StatusBadRequest)
	})

	t.Run("success - definition not found", func(t *testing.T) {
		mockInteractor.On("Get", "a random term").Return(&model.List{}, model.ErrNotFound{Term: "a random term"})
		controller := NewDefinitionController(mockInteractor)
		req, err := http.NewRequest(http.MethodGet, "/definitions/", nil)
		assert.Nil(t, err, "new request error should be nil")
		req = mux.SetURLVars(req, map[string]string{"term": "a random term"})
		rec := httptest.NewRecorder()
		controller.GetDefinitions(rec, req)
		assert.NotNil(t, rec.Body, "the body response should be not nil")
		assert.Equal(t, rec.Result().StatusCode, http.StatusNotFound)
	})

	t.Run("success - missing api key", func(t *testing.T) {
		mockInteractor.On("Get", "wfio").Return(&model.List{}, model.ErrMissingApiKey{})
		controller := NewDefinitionController(mockInteractor)
		req, err := http.NewRequest(http.MethodGet, "/definitions/", nil)
		assert.Nil(t, err, "new request error should be nil")
		req = mux.SetURLVars(req, map[string]string{"term": "wfio"})
		rec := httptest.NewRecorder()
		controller.GetDefinitions(rec, req)
		assert.NotNil(t, rec.Body, "the body response should be not nil")
		assert.Equal(t, rec.Result().StatusCode, http.StatusForbidden)
	})

}
func TestGetDefinitionsFromCSV(t *testing.T) {

	mockInteractor := MockInteractor{}

	t.Run("success - valid request by definition id", func(t *testing.T) {
		var list model.List
		mockInteractor.On("GetFromCSV", "1").Return(&model.List{
			Definitions: []*model.Definition{
				{Defid: 1},
			},
		}, nil)
		controller := NewDefinitionController(mockInteractor)
		req, err := http.NewRequest(http.MethodGet, "/definitions/csv/", nil)
		assert.Nil(t, err, "new request error should be nil")
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		rec := httptest.NewRecorder()
		controller.GetDefinitionsFromCSV(rec, req)
		res := rec.Result()
		assert.NotNil(t, res.Body, "the body response should be not nil")
		assert.Equal(t, res.StatusCode, http.StatusOK)
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &list)
		assert.Equal(t, list.Definitions[0].Defid, 1)
	})

	t.Run("success - definition not found", func(t *testing.T) {
		mockInteractor.On("GetFromCSV", "1233").Return(&model.List{}, model.ErrNotFoundInCSV{Id: "1233"})
		controller := NewDefinitionController(mockInteractor)
		req, err := http.NewRequest(http.MethodGet, "/definitions/csv/", nil)
		assert.Nil(t, err, "new request error should be nil")
		req = mux.SetURLVars(req, map[string]string{"id": "1233"})
		rec := httptest.NewRecorder()
		controller.GetDefinitionsFromCSV(rec, req)
		assert.NotNil(t, rec.Body, "the body response should be not nil")
		assert.Equal(t, rec.Result().StatusCode, http.StatusNotFound)
	})
}
