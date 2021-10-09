package presenter

import (
	"testing"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestResponseDefinitions(t *testing.T) {
	testcases := []struct {
		name              string
		response          *model.List
		expectedWrittenOn string
		parsedSuccessfull bool
	}{
		{
			name:              "success - written on field with correct format",
			expectedWrittenOn: "Sun, 30-April-2006 20:18",
			response: &model.List{
				Definitions: []*model.Definition{
					{
						Word:      "the",
						WrittenOn: "2006-04-30T20:18:42.000Z",
					},
				},
			},
			parsedSuccessfull: true,
		},
		{
			name:              "failure - written on field with incorrect format",
			expectedWrittenOn: "26-April-2018 Thursday",
			response: &model.List{
				Definitions: []*model.Definition{
					{
						Word:      "lmao",
						WrittenOn: "2018-04-26T19:19:47.118Z",
					},
				},
			},
			parsedSuccessfull: false,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			presenter := NewDefinitionPresenter()
			response := presenter.ResponseDefinitions(test.response)
			definition := response.Definitions[0]

			if test.parsedSuccessfull {
				assert.Equal(t, test.expectedWrittenOn, definition.WrittenOn)
			}

			if !test.parsedSuccessfull {
				assert.NotEqual(t, test.expectedWrittenOn, definition.WrittenOn)
			}
		})
	}
}
