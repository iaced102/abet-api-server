package service

import (
	"aBet/model"
	"aBet/usecase/repository"
)

type GetDetailReportService interface {
	GetDetailReport(reportId []string) (map[string][]model.DetailReport, error)
}

type getDetailReportService struct {
	detailReportRepository repository.GetDetailReportRepository
}

func NewGetDetailReportService(gDRR repository.GetDetailReportRepository) GetDetailReportService {
	return &getDetailReportService{
		detailReportRepository: gDRR,
	}
}

func (gDRS *getDetailReportService) GetDetailReport(reportId []string) (map[string][]model.DetailReport, error) {
	result := map[string][]model.DetailReport{}
	for _, s := range reportId {
		detailReport := model.DetailReport{

			ReportId: s,
		}
		subResult, e := gDRS.detailReportRepository.GetDetailReport(detailReport)
		if e != nil {
			return nil, e
		}
		result[s] = []model.DetailReport{}
		for _, i := range subResult {
			result[s] = append(result[s], i)
		}

	}
	return result, nil
}
