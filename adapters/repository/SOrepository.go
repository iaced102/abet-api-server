package repository

import "aBet/model"

type sORepository struct {
	db *Orm
}

type SORepsoitory interface {
	CreateSO(so *model.SODocument) error
	GetAllSODocument(*[]model.SODocument) error
	GetDetailSODocument(*model.SODocument) error
}

func NewSORepository(db *Orm) SORepsoitory {
	return &sORepository{
		db: db,
	}
}

func (sR *sORepository) CreateSO(so *model.SODocument) error {
	_, e := sR.db.pgdb.Model(so).Insert()
	return e
}

func (sR *sORepository) GetAllSODocument(so *[]model.SODocument) error {
	e := sR.db.pgdb.Model(so).Select()
	return e
}

func (sR *sORepository) GetDetailSODocument(so *model.SODocument) error {
	e := sR.db.pgdb.Model(so).Where("id = ?", so.Id).Select()
	return e
}
