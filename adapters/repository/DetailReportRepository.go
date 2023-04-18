package repository

import "aBet/model"

type DetailReportRepository interface {
	CreateDetailReport(dR model.DetailReport) error
	GetDetailReport(dR model.DetailReport) ([]model.DetailReport, error)
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

func (dRR *detailReportRepository) GetDetailReport(dR model.DetailReport) ([]model.DetailReport, error) {
	result := []model.DetailReport{}
	e := dRR.db.pgdb.Model(&result).Where("report_id = ?", dR.Id).Select()
	return result, e
}
