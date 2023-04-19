package service

import (
	"aBet/model"
	"aBet/usecase/repository"
)

type GetDetailReportService interface {
	GetDetailReport(reportId string) ([]model.DetailReport, error)
}

type getDetailReportService struct {
	detailReportRepository repository.GetDetailReportRepository
}

func NewGetDetailReportService(gDRR repository.GetDetailReportRepository) GetDetailReportService {
	return &getDetailReportService{
		detailReportRepository: gDRR,
	}
}

func (gDRS *getDetailReportService) GetDetailReport(reportId string) ([]model.DetailReport, error) {
	result, e := gDRS.detailReportRepository.GetDetailReport(reportId)
	if e != nil {
		return nil, e
	}

	return result, nil
}
