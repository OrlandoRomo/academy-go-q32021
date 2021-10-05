package registry

import (
	"github.com/OrlandoRomo/academy-go-q32021/interface/controller"
)

type registry struct {
	UrbanApiKey string
}
type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(apiKey string) Registry {
	return &registry{apiKey}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		List: r.NewListController(),
	}
}
