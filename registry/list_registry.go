package registry

import (
	"github.com/OrlandoRomo/academy-go-q32021/interface/controller"
	p "github.com/OrlandoRomo/academy-go-q32021/interface/presenter"
	re "github.com/OrlandoRomo/academy-go-q32021/interface/repository"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/interactor"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/presenter"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/repository"
)

func (r *registry) NewListController() controller.ListController {
	return controller.NewListController(r.NewListInteractor())
}

func (r *registry) NewListInteractor() interactor.ListInteractor {
	return interactor.NewListInteractor(r.NewListRepository(), r.NewListPresenter())
}

func (r *registry) NewListRepository() repository.ListRepository {
	return re.NewListRepository(r.UrbanApiKey)
}

func (r *registry) NewListPresenter() presenter.ListPresenter {
	return p.NewListPresenter()
}
