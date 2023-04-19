package service

import (
	"aBet/adapters/repository"
	"aBet/model"
	"time"
)

type CreateDocumentService interface {
	CreateDocument(userId string, name string, EvaluateField string, assessorId []string, verifierId []string, superviserId string) (model.Document, error)
}

type createDocumentService struct {
	documentRepository repository.DocumentRepository
}

func NewCreateDocumentService(rR repository.DocumentRepository) CreateDocumentService {
	return &createDocumentService{
		documentRepository: rR,
	}
}

func (cDS *createDocumentService) CreateDocument(userId string, name string, EvaluateField string, assessorId []string, verifierId []string, superviserId string) (model.Document, error) {
	doc := model.Document{
		EvaluteField: EvaluateField,
		CreatedBy:    userId,
		Name:         name,
		AssessorId:   assessorId,
		VerifierId:   verifierId,
		SuperviserId: superviserId,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		Id:           model.CreateUuid(),
	}
	e := cDS.documentRepository.CreateDocument(&doc)
	return doc, e
}
