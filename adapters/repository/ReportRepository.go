package repository

import "aBet/model"

type reportRepository struct {
	db *Orm
}
type ReportRepository interface {
	CreateReport(*model.Report) error
}

func NewReportRepository(db *Orm) ReportRepository {
	return &reportRepository{
		db: db,
	}
}

func (rR *reportRepository) CreateReport(report *model.Report) error {
	_, e := rR.db.pgdb.Model(report).Insert()
	return e
}
