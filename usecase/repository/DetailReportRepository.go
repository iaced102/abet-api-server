package repository

import "aBet/model"

type CreateDetailReportRepository interface {
	CreateDetailReport(dR model.DetailReport) error
}
