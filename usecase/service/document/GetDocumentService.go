package service

import (
	"aBet/model"
	"aBet/usecase/repository"
)

type GetDocumentService interface {
	GetDocument(documentId string) (model.Document, error)
}

type getDocumentService struct {
	documentRepository repository.GetDocumentRepository
}

func NewGetDocumentService(rR repository.GetDocumentRepository) GetDocumentService {
	return &getDocumentService{
		documentRepository: rR,
	}
}

func (cDS *getDocumentService) GetDocument(documentId string) (model.Document, error) {
	doc := model.Document{
		Id: documentId,
	}
	e := cDS.documentRepository.GetDocument(&doc)
	return doc, e
}
