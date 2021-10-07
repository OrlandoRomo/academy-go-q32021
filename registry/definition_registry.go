package registry

import (
	"github.com/OrlandoRomo/academy-go-q32021/interface/controller"
	re "github.com/OrlandoRomo/academy-go-q32021/interface/repository"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/interactor"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/repository"
)

func (r *registry) NewDefinitionController() controller.DefinitionController {
	return controller.NewDefinitionController(r.NewDefinitionInteractor())
}

func (r *registry) NewDefinitionInteractor() interactor.DefinitionInteractor {
	return interactor.NewDefinitionInteractor(r.NewDefinitionRepository())
}

func (r *registry) NewDefinitionRepository() repository.UrbanDictionaryRepository {
	return re.NewUrbanDictionaryRepository(r.UrbanDictionaryClient)
}
