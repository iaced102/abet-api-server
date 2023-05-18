package repository

import (
	"aBet/model"
	"fmt"
)

type DetailReportRepository interface {
	CreateDetailReport(dR model.DetailReport) error
	GetDetailReport(reportId string) ([]model.DetailReport, error)
	EditDetailReport(model.DetailReport) error
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

func (dRR *detailReportRepository) GetDetailReport(reportId string) ([]model.DetailReport, error) {
	result := []model.DetailReport{}
	e := dRR.db.pgdb.Model(&result).Where("report_id = ?", reportId).Select()
	return result, e
}

func (dRR *detailReportRepository) EditDetailReport(dR model.DetailReport) error {

	_, e := dRR.db.pgdb.Model(&dR).Where("id = ?", dR.Id).Update()
	fmt.Println(dR.Id, dR.Value)
	fmt.Println("__________________")
	fmt.Println(e)
	return e
}
