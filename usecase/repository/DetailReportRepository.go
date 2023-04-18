package repository

import "aBet/model"

type CreateDetailReportRepository interface {
	CreateDetailReport(dR model.DetailReport) error
}

type GetDetailReportRepository interface {
	GetDetailReport(dR model.DetailReport) ([]model.DetailReport, error)
}
