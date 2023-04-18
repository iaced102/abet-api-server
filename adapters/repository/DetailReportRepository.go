package repository

import "aBet/model"

type DetailReportRepository interface {
	CreateDetailReport(dR model.DetailReport) error
}

type detailReportRepository struct {
	db *Orm
}

func NewDetailReportRepository(db *Orm) DetailReportRepository {
	return &detailReportRepository{
		db: db,
	}
}

func (dRR *detailReportRepository) CreateDetailReport(dR model.DetailReport) error {
	_, e := dRR.db.pgdb.Model(&dR).Insert()
	return e
}
