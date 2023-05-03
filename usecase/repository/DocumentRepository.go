package repository

import "aBet/model"

type CreateDocumentRepository interface{}

type GetDocumentRepository interface {
	GetDocument(doc *model.Document) error
	GetAllDocument(doc *[]model.Document) error
	GetAllDocumentByUserId(userId string) ([]model.Document, error)
	GetAllPIbySOId(sOId string) ([]model.Document, error)
}
