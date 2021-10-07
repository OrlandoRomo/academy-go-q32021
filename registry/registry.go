package registry

import (
	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/api"
	"github.com/OrlandoRomo/academy-go-q32021/interface/controller"
)

type registry struct {
	UrbanDictionaryClient *api.UrbanDictionary
}
type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(urbanDictionary *api.UrbanDictionary) Registry {
	return &registry{urbanDictionary}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Definitions: r.NewDefinitionController(),
	}
}
