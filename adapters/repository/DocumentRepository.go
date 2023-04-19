package repository

import (
	"aBet/model"
	"fmt"
)

type documentRepository struct {
	db *Orm
}
type DocumentRepository interface {
	CreateDocument(doc *model.Document) error
	GetDocument(doc *model.Document) error
	GetAllDocument(doc *[]model.Document) error
	GetAllDocumentByUserId(userId string) ([]model.Document, error)
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

func (dR *documentRepository) GetAllDocument(doc *[]model.Document) error {
	e := dR.db.pgdb.Model(doc).Select()
	return e
}

func (dR *documentRepository) GetAllDocumentByUserId(userId string) ([]model.Document, error) {
	doc := []model.Document{}
	fmt.Println("created_by = '" + userId + "' or assessor_id like '%" + userId + "%' or verifier_id like '%" + userId + "%' or superviser_id = '" + userId + "'")
	e := dR.db.pgdb.Model(&doc).Where("created_by = '" + userId + "' or assessor_id like '%" + userId + "%' or verifier_id like '%" + userId + "%' or superviser_id = '" + userId + "'").Select()
	return doc, e
}
