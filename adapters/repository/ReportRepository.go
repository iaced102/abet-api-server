package repository

import "aBet/model"

type reportRepository struct {
	db *Orm
}
type ReportRepository interface {
	CreateReport(*model.Report) error
	GetAllReport(*model.Report) ([]model.Report, error)
	DeleteDocument(id string) error
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

func (rR *reportRepository) GetAllReport(report *model.Report) ([]model.Report, error) {

	result := []model.Report{}
	e := rR.db.pgdb.Model(&result).Where("document_id = ?", report.DocumentId).Select()

	return result, e
}

func (rR *reportRepository) DeleteDocument(id string) error {
	do := model.Document{Id: id}
	_, e := rR.db.pgdb.Model(&do).Where("id = ?", do.Id).Delete()
	return e
}
