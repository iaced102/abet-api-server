package service

import (
	"aBet/model"
	"aBet/usecase/repository"
)

type GetDocumentService interface {
	GetDocument(documentId string) (model.Document, error)
	GetAllDocument() ([]model.Document, error)
	GetAllDocumentByUserId(userId string) ([]model.Document, error)
	GetAllPIbySOId(userId string) ([]model.Document, error)
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

func (cDS *getDocumentService) GetAllDocument() ([]model.Document, error) {
	doc := []model.Document{}
	e := cDS.documentRepository.GetAllDocument(&doc)
	return doc, e
}
func (cDS *getDocumentService) GetAllDocumentByUserId(userId string) ([]model.Document, error) {

	doc, e := cDS.documentRepository.GetAllDocumentByUserId(userId)
	return doc, e
}

func (cDS *getDocumentService) GetAllPIbySOId(sOId string) ([]model.Document, error) {

	doc, e := cDS.documentRepository.GetAllDocumentByUserId(sOId)
	return doc, e
}
