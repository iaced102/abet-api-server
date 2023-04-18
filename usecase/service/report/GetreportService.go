package service

import (
	"aBet/model"
	"aBet/usecase/repository"
)

type GetReportService interface {
	GetAllReport(documentId string) ([]model.Report, error)
}

type getReportService struct {
	reportRepository repository.GetReportRepository
}

func NewGetReportService(rR repository.GetReportRepository) GetReportService {
	return &getReportService{
		reportRepository: rR,
	}
}

func (cR *getReportService) GetAllReport(documentId string) ([]model.Report, error) {

	report := model.Report{
		DocumentId: documentId,
	}
	result := []model.Report{}
	result, e := cR.reportRepository.GetAllReport(&report)
	if len(result) > 1 {
		for i, v := range result {
			if v.IsTemplate == 1 {
				result = append(result[:i], result[i+1:]...)
			}
		}
	}
	return result, e
}
