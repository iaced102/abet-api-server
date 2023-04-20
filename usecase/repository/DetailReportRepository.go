package repository

import "aBet/model"

type CreateDetailReportRepository interface {
	CreateDetailReport(dR model.DetailReport) error
}

type GetDetailReportRepository interface {
	GetDetailReport(reportId string) ([]model.DetailReport, error)
}

type EditDetailReportRepository interface {
	EditDetailReport(model.DetailReport) error
}
