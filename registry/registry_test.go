package registry

import (
	"testing"

	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/api"
	"github.com/OrlandoRomo/academy-go-q32021/interface/controller"
	"github.com/stretchr/testify/assert"
)

var (
	testApiKey = "cc2464e4a1mshb5ceeca91e5a6adp1fa80bjsn4b48e2408b87"
)

func TestNewRegistry(t *testing.T) {
	testcases := []struct {
		name        string
		message     string
		urbanClient *api.UrbanDictionary
	}{
		{
			name:        "success - urban client set up",
			message:     "urban client in registry set up",
			urbanClient: api.NewUrbanDictionary(testApiKey),
		},
		{
			name:        "failure - urban client nil",
			message:     "registry app with a nil urban client",
			urbanClient: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			if test.urbanClient == nil {
				assert.Nil(t, test.urbanClient, test.message)
			}
			if test.urbanClient != nil {
				assert.NotNil(t, test.urbanClient, test.message)
			}
		})
	}
}

func TestNewAppController(t *testing.T) {
	appController := NewRegistry(nil).NewAppController()
	t.Run("success - return AppController type", func(t *testing.T) {
		assert.IsType(t, controller.AppController{}, appController, "app contoller type received")
	})

	t.Run("failure - return another type", func(t *testing.T) {
		assert.NotSame(t, controller.AppController{}, struct{}{}, "expected AppControll type")
	})
}
