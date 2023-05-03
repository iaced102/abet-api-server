package service

import (
	"aBet/model"
	"aBet/usecase/repository"
)

type GetAllSODocumentService interface {
	GetAllSO() ([]model.SODocument, error)
	GetDetailSODocument(id string) (model.SODocument, error)
}
type getAllSODocumentService struct {
	soRepository repository.GetSORepository
}

func NewGetAllSODocumentService(s repository.GetSORepository) GetAllSODocumentService {
	return &getAllSODocumentService{
		soRepository: s,
	}
}

func (gASS *getAllSODocumentService) GetAllSO() ([]model.SODocument, error) {
	lSO := []model.SODocument{}
	e := gASS.soRepository.GetAllSODocument(&lSO)
	return lSO, e
}

func (gASS *getAllSODocumentService) GetDetailSODocument(id string) (model.SODocument, error) {
	lSO := model.SODocument{
		Id: id,
	}
	e := gASS.soRepository.GetDetailSODocument(&lSO)
	return lSO, e
}
