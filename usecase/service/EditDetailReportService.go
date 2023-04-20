package service

import (
	"aBet/model"
	"aBet/usecase/repository"
)

type EditDetailReportService interface {
	EditDetailReport(model.DetailReport) error
}
type editDetailReportService struct {
	editDetailReportRepository repository.EditDetailReportRepository
}

func NewEditDetailReportService(eDRR repository.EditDetailReportRepository) EditDetailReportService {
	return &editDetailReportService{
		editDetailReportRepository: eDRR,
	}
}

func (eDRS *editDetailReportService) EditDetailReport(dR model.DetailReport) error {
	return eDRS.editDetailReportRepository.EditDetailReport(dR)
}
