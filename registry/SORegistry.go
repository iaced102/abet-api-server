package registry

import (
	"aBet/adapters/controller"
	"aBet/adapters/repository"
	"aBet/usecase/service"
)

func (r *registry) NewSOController() controller.SOController {
	return controller.NewSOController(r.NewCreateSOService(), r.NewGetSOService())
}

func (r *registry) NewCreateSOService() service.CreateSOService {
	return service.NewCreateSOService(r.NewSORepository())
}

func (r *registry) NewGetSOService() service.GetAllSODocumentService {
	return service.NewGetAllSODocumentService(r.NewSORepository())
}

func (r *registry) NewSORepository() repository.SORepsoitory {
	return repository.NewSORepository(r.db)
}
