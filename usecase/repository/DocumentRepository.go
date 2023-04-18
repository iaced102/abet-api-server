package repository

import "aBet/model"

type CreateDocumentRepository interface{}

type GetDocumentRepository interface {
	GetDocument(doc *model.Document) error
}
