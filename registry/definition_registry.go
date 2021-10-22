package registry

import (
	"github.com/OrlandoRomo/academy-go-q32021/interface/controller"
	p "github.com/OrlandoRomo/academy-go-q32021/interface/presenter"
	re "github.com/OrlandoRomo/academy-go-q32021/interface/repository"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/interactor"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/presenter"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/repository"
)

func (r *registry) NewDefinitionController() controller.DefinitionController {
	return controller.NewDefinitionController(r.NewDefinitionInteractor())
}

func (r *registry) NewDefinitionInteractor() interactor.DefinitionInteractor {
	return interactor.NewDefinitionInteractor(r.NewDefinitionRepository(), r.NewDefinitionPresenter())
}

func (r *registry) NewDefinitionRepository() repository.UrbanDictionaryRepository {
	return re.NewUrbanDictionaryRepository(r.UrbanDictionaryClient)
}

func (r *registry) NewDefinitionPresenter() presenter.DefinitionPresenter {
	return p.NewDefinitionPresenter()
}
