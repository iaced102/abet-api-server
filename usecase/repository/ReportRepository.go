package repository

import "aBet/model"

type ReportRepository interface {
	CreateReport(*model.Report) error
	DeleteDocument(id string) error
}

type GetReportRepository interface {
	GetAllReport(*model.Report) ([]model.Report, error)
}
