package service

import (
	"aBet/model"
	"aBet/usecase/repository"

	"github.com/google/uuid"
)

type CreateReportService interface {
	CreateReport(documentId string, field string, listStudent []model.Student) (model.Report, error)
}

type createReportService struct {
	reportRepository repository.ReportRepository
}

func NewCreateReportService(rR repository.ReportRepository) CreateReportService {
	return &createReportService{
		reportRepository: rR,
	}
}

func (cR *createReportService) CreateReport(documentId string, field string, listStudent []model.Student) (model.Report, error) {
	report := model.Report{
		DocumentId: documentId,
		Field:      field,
		Id:         uuid.NewString(),
	}
	e := cR.reportRepository.CreateReport(&report)
	return report, e
}
