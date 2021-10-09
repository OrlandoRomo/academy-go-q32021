package api

import (
	"testing"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/stretchr/testify/assert"
)

var (
	testApiKey = "cc2464e4a1mshb5ceeca91e5a6adp1fa80bjsn4b48e2408b87"
)

func TestUpdateCSVPath(t *testing.T) {
	t.Run("set valid CSV path", func(t *testing.T) {
		urbanClient := NewUrbanDictionary(testApiKey)
		urbanClient.UpdateCSVPath("/usr/local/definitions.csv")
		assert.Equal(t, urbanClient.CSVPath, "/usr/local/definitions.csv")
	})
}

func TestGetDefinitions(t *testing.T) {
	urbanClint := NewUrbanDictionary(testApiKey)
	t.Run("success - definition by term", func(t *testing.T) {
		definitions, err := urbanClint.GetDefinitions("ha")
		assert.Nil(t, err)
		assert.NotNil(t, definitions)
	})

	t.Run("success - definition not found", func(t *testing.T) {
		definitions, err := urbanClint.GetDefinitions("carnaval")
		assert.Equal(t, err, model.ErrNotFound{Term: "carnaval"})
		assert.Nil(t, definitions)
	})
}

func TestGetDefinitionById(t *testing.T) {
	urbanClint := NewUrbanDictionary(testApiKey)
	t.Run("success - definition found in CSV", func(t *testing.T) {
		urbanClint.UpdateCSVPath("../../data/definitions_test.csv")
		definitions, err := urbanClint.GetDefinitionById("10593002")
		assert.Nil(t, err)
		assert.NotNil(t, definitions)
	})

	t.Run("success - definition not found in CSV", func(t *testing.T) {
		urbanClint.UpdateCSVPath("../../data/definitions_test.csv")
		definitions, err := urbanClint.GetDefinitionById("1232")
		assert.EqualValues(t, model.ErrNotFoundInCSV{"1232"}, err)
		assert.Nil(t, definitions)
	})
}

func TestOpen(t *testing.T) {
	urbanClint := NewUrbanDictionary(testApiKey)
	t.Run("success - open CSV file", func(t *testing.T) {
		urbanClint.UpdateCSVPath("../../data/definitions_test.csv")
		_, err := urbanClint.Open()
		assert.Nil(t, err)
	})
}

func TestRead(t *testing.T) {
	urbanClint := NewUrbanDictionary(testApiKey)
	t.Run("success - reading content of the CSV", func(t *testing.T) {
		urbanClint.UpdateCSVPath("../../data/definitions_test.csv")
		definitions, err := urbanClint.Read("10593002")
		assert.Nil(t, err)
		assert.NotNil(t, definitions)
		assert.Equal(t, 1, len(definitions))
	})
}

func TestWrite(t *testing.T) {
	urbanClint := NewUrbanDictionary(testApiKey)
	t.Run("success - reading content of the CSV", func(t *testing.T) {
		urbanClint.UpdateCSVPath("../../data/definitions_test.csv")
		err := urbanClint.Write(&model.List{
			Definitions: []*model.Definition{
				{Defid: 12345, Word: "Hola"},
			},
		})
		assert.Nil(t, err)
	})
}
