package repository

import "aBet/model"

type documentRepository struct {
	db *Orm
}
type DocumentRepository interface {
	CreateDocument(doc *model.Document) error
	GetDocument(doc *model.Document) error
}

func NewDocumentRepository(db *Orm) DocumentRepository {
	return &documentRepository{
		db: db,
	}
}

func (dR *documentRepository) CreateDocument(doc *model.Document) error {
	_, e := dR.db.pgdb.Model(doc).Insert()
	return e
}

func (dR *documentRepository) GetDocument(doc *model.Document) error {
	e := dR.db.pgdb.Model(doc).Where("id = ?", doc.Id).Select()
	return e
}
