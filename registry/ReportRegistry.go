package registry

import (
	"aBet/adapters/controller"
	"aBet/adapters/repository"
	"aBet/usecase/service"
	dservice "aBet/usecase/service/document"
	rService "aBet/usecase/service/report"
)

func (r *registry) NewReportController() controller.ReportController {
	return controller.NewReportController(r.NewCreateReportService(), r.NewCreateDocumentService(), r.NewCreateDetailReportService())
}

func (r *registry) NewCreateReportService() rService.CreateReportService {
	return rService.NewCreateReportService(r.NewReportRepository())
}
func (r *registry) NewReportRepository() repository.ReportRepository {
	return repository.NewReportRepository(r.db)
}

func (r *registry) NewCreateDocumentService() dservice.CreateDocumentService {
	return dservice.NewCreateDocumentService(r.NewDocumentRepository())
}
func (r *registry) NewDocumentRepository() repository.DocumentRepository {
	return repository.NewDocumentRepository(r.db)
}

func (r *registry) NewCreateDetailReportService() service.CreateDetailReportService {
	return service.NewCreateDetailReportService(r.NewCreateDetailReportRepository())
}
func (r *registry) NewCreateDetailReportRepository() repository.DetailReportRepository {
	return repository.NewDetailReportRepository(r.db)
}
