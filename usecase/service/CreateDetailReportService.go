package service

import (
	"aBet/model"
	"aBet/usecase/repository"

	"github.com/google/uuid"
)

type CreateDetailReportService interface {
	CreateListDetailReport(listStudent []model.Student, reportId string, field string, template string) error
}

type createDetailReportService struct {
	detailReportRepository repository.CreateDetailReportRepository
}

func NewCreateDetailReportService(rR repository.CreateDetailReportRepository) CreateDetailReportService {
	return &createDetailReportService{
		detailReportRepository: rR,
	}
}

func (cDRS *createDetailReportService) CreateListDetailReport(listStudent []model.Student, reportId string, field string, template string) error {
	for _, s := range listStudent {
		detailReport := model.DetailReport{
			Id:        uuid.NewString(),
			StudentId: s.StudentId,
			FirstName: s.FirstName,
			LastName:  s.LastName,
			ClassId:   s.ClassId,
			ReportId:  reportId,
			Value:     template,
		}
		e := cDRS.detailReportRepository.CreateDetailReport(detailReport)
		if e != nil {
			return e
		}
	}
	return nil
}
